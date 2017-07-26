using System;
using System.Collections.Generic;
using System.Linq;
using SQLite;
using Xamarin.Forms;

namespace App1
{
    public class GooglePlacesDatabase
    {
        private static GooglePlacesDatabase selfInstance = null;
        public static GooglePlacesDatabase getInstance
        {
            get
            {
                if (selfInstance == null) selfInstance = new GooglePlacesDatabase();
                return selfInstance;
            }
        }

        static object mlock = new object();
        SQLiteConnection sqLiteConnection;

        public GooglePlacesDatabase()
        {
            sqLiteConnection = DependencyService.Get<ISqlite>().GetConnection();
            sqLiteConnection.CreateTable<GooglePlacesItems>();
        }

        public IEnumerable<GooglePlacesItems> GetItems()
        {
            lock (mlock)
            {
                return (from i in sqLiteConnection.Table<GooglePlacesItems>() select i).ToList();
            }
        }

        public IEnumerable<GooglePlacesItems> GetSearchAddress(string sSearch)
        {
            lock (mlock)
            {
                return sqLiteConnection.Query<GooglePlacesItems>(String.Format("SELECT * FROM [tbPlaceId] WHERE [Address] LIKE \"%{0}%\"", sSearch));
            }
        }

        public IEnumerable<GooglePlacesItems> GetSearchPcName(string sSearch)
        {
            lock (mlock)
            {
                return sqLiteConnection.Query<GooglePlacesItems>(String.Format("SELECT * FROM [tbPlaceId] WHERE [Name] LIKE \"%{0}%\"", sSearch));
            }
        }
    }
}
