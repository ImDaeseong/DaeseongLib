using System.IO;
using Xamarin.Forms;
using SQLite;
using App1.Droid;

[assembly: Dependency(typeof(CSqlite))]
namespace App1.Droid
{
    public class CSqlite : ISqlite
    {
        public CSqlite()
        {
        }

        public SQLiteConnection GetConnection()
        {
            var sqliteFilename = "GooglePlaces.db";
            string documentsPath = System.Environment.GetFolderPath(System.Environment.SpecialFolder.Personal);
            var path = Path.Combine(documentsPath, sqliteFilename);

            if (!File.Exists(path))
            {
                CopyDatabaseIfNotExists(path, sqliteFilename);
            }

            var conn = new SQLite.SQLiteConnection(path);
            return conn;
        }

        private static void CopyDatabaseIfNotExists(string dbPath, string sqliteFilename)
        {
            using (var br = new BinaryReader(Android.App.Application.Context.Assets.Open(sqliteFilename)))
            {
                using (var bw = new BinaryWriter(new FileStream(dbPath, FileMode.Create)))
                {
                    byte[] buffer = new byte[2048];
                    int length = 0;
                    while ((length = br.Read(buffer, 0, buffer.Length)) > 0)
                    {
                        bw.Write(buffer, 0, length);
                    }
                }
            }
        }

    }
}