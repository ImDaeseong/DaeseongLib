using System;
using System.Data.SQLite;
using System.IO;

namespace WindowsFormsApplication1
{
    class clsDBConnect
    {
        private object m_Lock = new object();
        private SQLiteConnection sqlConn = new SQLiteConnection();
        clsComm CFunc = clsComm.getInstance;

        private static clsDBConnect selfInstance = null;
        public static clsDBConnect getInstance
        {
            get
            {
                if (selfInstance == null) selfInstance = new clsDBConnect();
                return selfInstance;
            }
        }

        public bool DisConnect()
        {
            bool bResult = true;
            try
            {
                sqlConn.Close();
                sqlConn.Dispose();
            }
            catch
            {
                bResult = false;
            }
            return bResult;
        }

        public bool DbConnection()
        {
            bool bConnected = true;

            try
            {
                String databaseName = "GooglePlaces.db";
                if (File.Exists(databaseName))
                {
                    sqlConn = new SQLiteConnection("Data Source=" + databaseName);
                    sqlConn.Open();
                }
                else
                {
                    SQLiteConnection.CreateFile(databaseName);
                    sqlConn = new SQLiteConnection("Data Source=" + databaseName);
                    sqlConn.Open();
                }
            }
            catch
            {
                bConnected = false;
            }
            return bConnected;
        }

        public void ExeCuteQuery(string sQuery)
        {
            lock (m_Lock)
            {
                SQLiteCommand cmd = new SQLiteCommand(sQuery, sqlConn);
                cmd.ExecuteNonQuery();
                cmd.Dispose();
            }
        }

        public SQLiteDataReader GetSelectQuery(string sQuery)
        {
            SQLiteDataReader r = null;
            lock (m_Lock)
            {
                SQLiteCommand cmd = new SQLiteCommand(sQuery, sqlConn);
                r = cmd.ExecuteReader();
                cmd.Dispose();
            }
            return r;
        }

        public bool createTableGooglePlaces()
        {
            bool bCreated = true;
            try
            {
                string sQuery = "CREATE TABLE tbPlaceId ( " +
                    "rIndex INTEGER PRIMARY KEY AUTOINCREMENT," +
                    "place_id varchar(100));";
                SQLiteCommand cmd = new SQLiteCommand(sQuery, sqlConn);
                cmd.ExecuteNonQuery();
            }
            catch 
            {
                bCreated = false;
            }
            return bCreated;
        }

        public void ReadDataFile()
        {
            string sQuery = string.Format("SELECT * FROM tbPlaceId");
            SQLiteDataReader r = GetSelectQuery(sQuery);
            while (r.Read())
            {               
                string place_id = string.Format("{0}", r["place_id"].ToString());
                CFunc.GetMainForm().SetLogText("ReadDataFile:" + place_id);
            }
        }

    }
}
