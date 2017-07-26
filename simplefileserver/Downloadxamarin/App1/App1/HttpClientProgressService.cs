using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using Xamarin.Forms;

namespace App1
{
    public class HttpClientProgressService
    {
        private IDownloadPath dPath;
        private HttpClient client;
        private ProgressBar progress;

        public HttpClientProgressService(ProgressBar pg)
        {
            client = new HttpClient();
            dPath = DependencyService.Get<IDownloadPath>();
            progress = pg;
        }

        private string GetFileName(string strFilename)
        {
            int nPos = strFilename.LastIndexOf('/');
            int nLength = strFilename.Length;
            if (nPos < nLength)
                return strFilename.Substring(nPos + 1, (nLength - nPos) - 1);
            return strFilename;
        }

        public async Task DownloadFileAsync(string url, CancellationToken token, Action<string> onFail, Action onComplete)
        {
            try
            {
                Uri uri = new Uri(url);
                string filename = System.IO.Path.GetFileName(uri.LocalPath);
                Stream fileStream = null;

                try
                {
                    fileStream = dPath.GetFileStream(filename);
                }
                catch
                {
                    onFail($"Failed to get file writer");
                    return;
                }

                var response = await client.GetAsync(url, HttpCompletionOption.ResponseHeadersRead, token);

                if (!response.IsSuccessStatusCode)
                {
                    onFail($"Failed to download file, Status code : {response.StatusCode}");
                    return;
                }

                var total = response.Content.Headers.ContentLength.HasValue ? response.Content.Headers.ContentLength.Value : -1L;
                var canReportProgress = total != -1 && progress != null;

                using (var stream = await response.Content.ReadAsStreamAsync())
                {
                    var totalRead = 0L;
                    var buffer = new byte[4096];
                    var isMoreToRead = true;

                    do
                    {
                        token.ThrowIfCancellationRequested();

                        var read = await stream.ReadAsync(buffer, 0, buffer.Length, token);

                        if (read == 0)
                        {
                            isMoreToRead = false;
                        }
                        else
                        {
                            var data = new byte[read];
                            buffer.ToList().CopyTo(0, data, 0, read);

                            fileStream.Write(data, 0, read);
                                                        
                            totalRead += read;

                            if (canReportProgress)
                            {
                                progress.Progress = (totalRead * 1d) / (total * 1d) * 100;
                            }
                        }
                    } while (isMoreToRead);

                    onComplete();
                }
            }
            catch (Exception ex)
            {
                Debug.WriteLine(@"Exception {0}", ex.Message);
                onFail($"Failed to download file : {ex.Message}");
            }
        }

    }
}
