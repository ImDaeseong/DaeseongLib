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

        public SQLiteConnection GetSQLiteConnection
        {
            get
            {
                return sqlConn;
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
                string sQuery = "CREATE TABLE tbPlaceId  ( " +
                    "rIndex INTEGER PRIMARY KEY AUTOINCREMENT," +
                    "Name varchar(100)," +
                    "Lat varchar(20)," +
                    "Lng varchar(20)," +
                    "place_id varchar(100)," +
                    "Address varchar(200)," +
                    "phonenumber varchar(30)," +
                    "Internatphonenumber varchar(30)," +
                    "Url varchar(100)," +
                    "Authorname varchar(100)," +
                    "Authorurl varchar(100)," +
                    "Profilephotourl varchar(100)," +
                    "Relativetimedescription varchar(50)," +
                    "TextDesc varchar(500));";
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
                String s1 = r["rIndex"].ToString();
                String s2 = r["Name"].ToString();
                String s3 = r["Lat"].ToString();
                String s4 = r["Lng"].ToString();
                String s5 = r["place_id"].ToString();
                String s6 = r["Address"].ToString();
                String s7 = r["phonenumber"].ToString();
                String s8 = r["Internatphonenumber"].ToString();
                String s9 = r["Url"].ToString();
                String s10 = r["Authorname"].ToString();
                String s11 = r["Authorurl"].ToString();
                String s12 = r["Profilephotourl"].ToString();
                String s13 = r["Relativetimedescription"].ToString();
                String s14 = r["TextDesc"].ToString();

                string sMsg = string.Format("{0}|{1}|{2}|{3}|{4}|{5}|{6}|{7}|{9}|{10}|{11}|{12}|{13}", s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14);
            }
        }        

    }
}
