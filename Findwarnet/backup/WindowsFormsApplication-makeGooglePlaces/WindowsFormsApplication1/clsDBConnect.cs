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
                string sQuery = "CREATE TABLE FormattedAddress ( " +
                    "rIndex INTEGER PRIMARY KEY AUTOINCREMENT," +
                    "Name varchar(100)," +
                    "Lat varchar(20)," +
                    "Lng varchar(20)," +
                    "place_id varchar(100)," +
                    "Address varchar(200));";
                SQLiteCommand cmd = new SQLiteCommand(sQuery, sqlConn);
                cmd.ExecuteNonQuery();
            }
            catch 
            {
                bCreated = false;
            }
            return bCreated;
        }
                
        private void ReadDataFile()
        {
            string sQuery = string.Format("SELECT * FROM FormattedAddress");
            SQLiteDataReader r = GetSelectQuery(sQuery);
            while (r.Read())
            {
                String s1 = r["rIndex"].ToString();
                String s2 = r["Name"].ToString();
                String s3 = r["Lat"].ToString();
                String s4 = r["Lng"].ToString();
                String s5 = r["place_id"].ToString();
                String s6 = r["Address"].ToString();
                string val = string.Format("{0}|{1}|{2}|{3}|{4}|{5}", s1, s2, s3, s4, s5, s6);
                Console.WriteLine(val);
            }
        }

        static readonly object mlockWrite = new object();

        private void WriteDataFile()
        {
            try
            {
                for (int i = 1; i < 19; i++)
                {
                    string sPath = string.Format("Mapinfo_{0}.txt", i);
                    StreamReader reader = new StreamReader(sPath);
                    using (reader)
                    {
                        string Name;
                        string Lat;
                        string Lng;
                        string placeid;
                        string Address;
                        
                        string line = reader.ReadLine();
                        while (line != null)
                        {
                            string[] value = line.Split('|');

                            try
                            {

                                Name = CFunc.NullVal(value[0], "").Replace("\"", "").Replace("\r", "").Replace("\n", "").Replace("\t", "").Trim();
                            }
                            catch { Name = ""; }

                            try
                            {
                                Lat = CFunc.NullVal(value[1], "").Replace("\"", "").Replace("\r", "").Replace("\n", "").Replace("\t", "").Trim();
                            }
                            catch { Lat = ""; }

                            try
                            {
                                Lng = CFunc.NullVal(value[2], "").Replace("\"", "").Replace("\r", "").Replace("\n", "").Replace("\t", "").Trim();
                            }
                            catch { Lng = ""; }

                            try
                            {
                                placeid = CFunc.NullVal(value[3], "").Replace("\"", "").Replace("\r", "").Replace("\n", "").Replace("\t", "").Trim();
                            }
                            catch { placeid = ""; }

                            try
                            {
                                Address = CFunc.NullVal(value[4], "").Replace("\"", "").Replace("\r", "").Replace("\n", "").Replace("\t", "").Trim();
                            }
                            catch { Address = ""; }

                            lock (mlockWrite)
                            {
                                string sQuery = string.Format("SELECT * FROM FormattedAddress where Lat={0} and Lng={1}", CFunc.QStr(Lat), CFunc.QStr(Lng));
                                SQLiteDataReader r = GetSelectQuery(sQuery);
                                if (!r.Read())
                                {
                                    sQuery = "";
                                    sQuery = string.Format("INSERT INTO FormattedAddress ( Name, Lat, Lng, place_id, Address) VALUES ({0},{1},{2},{3},{4});", CFunc.QStr(Name), CFunc.QStr(Lat), CFunc.QStr(Lng), CFunc.QStr(placeid), CFunc.QStr(Address));
                                    ExeCuteQuery(sQuery);
                                }
                            }
                            line = reader.ReadLine();
                        }
                    }
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());
            }
        }

    }
}
