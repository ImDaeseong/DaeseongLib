using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using System.Diagnostics;
using System.IO;
using Xamarin.Forms;

namespace App1
{
    public class HttpClientService
    {
        private IDownloadPath dPath;
        private HttpClient client;

        public HttpClientService()
        {
            client = new HttpClient();
            dPath = DependencyService.Get<IDownloadPath>();
        }       
  
        private string GetFileName(string strFilename)
        {
            int nPos = strFilename.LastIndexOf('/');
            int nLength = strFilename.Length;
            if (nPos < nLength)
                return strFilename.Substring(nPos + 1, (nLength - nPos) - 1);
            return strFilename;
        }

        public async Task<string> DownloadTextAsync(string sUrl)
        {
            string sReturn = "";

            var uri = new Uri(string.Format(sUrl, string.Empty));

            try
            {
                var response = await client.GetAsync(uri);
                if (response.IsSuccessStatusCode)
                {
                    var content = await response.Content.ReadAsStringAsync();
                    sReturn = content;
                }
            }
            catch (Exception ex)
            {
                Debug.WriteLine(@"Exception {0}", ex.Message);
            }
            return sReturn;
        }

        public async Task<bool> DownloadFileAsync(string sUrl)
        {
            bool bDownload = false;        
                      
            var uri = new Uri(string.Format(sUrl, string.Empty));
            
            try
            {
                var response = await client.GetAsync(uri);
                if (response.IsSuccessStatusCode)
                {
                    string slocalPath = string.Format("{0}", dPath.GetFolderPath(GetFileName(sUrl)));
                    
                    var content = await response.Content.ReadAsByteArrayAsync();
                    dPath.WriteFile(slocalPath, content);

                    bDownload = true;
                }
            }
            catch (Exception ex)
            {
                Debug.WriteLine(@"Exception {0}", ex.Message);
                bDownload = false;
            }

            return bDownload;
        }

    }
}
