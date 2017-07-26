using System;
using Xamarin.Forms;
using UIKit;
using App1.iOS;
using System.IO;
using System.Diagnostics;

[assembly: Xamarin.Forms.Dependency(typeof(CDownloadPath))]
namespace App1.iOS
{
    public class CDownloadPath : IDownloadPath
    {       
        public string GetFolderPath(string fileName)
        {
            return "";           
            //string path = Android.OS.Environment.ExternalStorageDirectory.AbsolutePath;
            //return System.IO.Path.Combine(path, Android.OS.Environment.DirectoryDownloads, fileName);
        }

        public bool WriteFile(string slocalPath, byte[] content)
        {
            bool bCreate = false;
            try
            {
                File.WriteAllBytes(slocalPath, content);
                bCreate = true;
            }
            catch (Exception ex)
            {
                Debug.WriteLine(@"Exception {0}", ex.Message);
                bCreate = false;
            }
            return bCreate;
        }

        public Stream GetFileStream(string fileName)
        {
            return null;
            /*
            string path = Android.OS.Environment.ExternalStorageDirectory.AbsolutePath;
            var sFullPath = System.IO.Path.Combine(path, Android.OS.Environment.DirectoryDownloads, fileName);

            try
            {
                if (File.Exists(sFullPath))
                    File.Delete(sFullPath);

                return new FileStream(sFullPath, FileMode.CreateNew);
            }
            catch (Exception ex)
            {
                Debug.WriteLine(@"Exception {0}", ex.Message);
            }
            return null;
            */
        }

    }
}