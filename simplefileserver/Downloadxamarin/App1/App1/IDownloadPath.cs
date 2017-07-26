using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;


namespace App1
{
    public interface IDownloadPath
    {
        string GetFolderPath(string fileName);

        bool WriteFile(string slocalPath, byte[] content);

        Stream GetFileStream(string fileName);
    }
}
