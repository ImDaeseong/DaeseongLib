using System;
using System.IO;
using Foundation;
using Xamarin.Forms;
using SQLite;
using App1.iOS;

[assembly: Dependency(typeof(CSqlite))]
namespace App1.iOS
{
    public class CSqlite : ISqlite
    {
        public CSqlite()
        {
        }

        public SQLiteConnection GetConnection()
        {
            var sqliteFilename = "GooglePlaces.db";
            string documentsPath = Environment.GetFolderPath(Environment.SpecialFolder.Personal);
            string libraryPath = Path.Combine(documentsPath, "..", "Library");
            var path = Path.Combine(libraryPath, sqliteFilename);

            if (!File.Exists(path))
            {
                var existingDb = NSBundle.MainBundle.PathForResource("daeseong", "db");
                File.Copy(existingDb, path);
            }

            var conn = new SQLite.SQLiteConnection(path);
            return conn;
        }
    }
}