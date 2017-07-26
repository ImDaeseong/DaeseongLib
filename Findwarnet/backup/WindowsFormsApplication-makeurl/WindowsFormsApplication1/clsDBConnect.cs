using System.Data;
using System.Data.SqlClient;

namespace WindowsFormsApplication1
{
    class clsDBConnect
    {        
        private object m_Lock = new object();
        
        private static clsDBConnect selfInstance = null;
        public static clsDBConnect getInstance
        {
            get
            {
                if (selfInstance == null) selfInstance = new clsDBConnect();
                return selfInstance;
            }
        }

        private SqlConnection sqlConn = new SqlConnection();
                
        public int ExeCuteQuery(string sQuery)
        {
            int nRtn = 0;
            lock (m_Lock)
            {
                SqlCommand cmd = new SqlCommand();
                cmd.CommandText = sQuery;
                cmd.Connection = sqlConn;
                nRtn = cmd.ExecuteNonQuery();
                cmd.Dispose();
            }
            return nRtn;
        }
        
        public DataSet GetSelectQuery(string sQuery, DataSet ds)
        {
            lock (m_Lock)
            {
                SqlDataAdapter adt = new SqlDataAdapter(sQuery, sqlConn);
                adt.Fill(ds);
                adt.Dispose();
            }
            return ds;
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

            sqlConn.ConnectionString = @"Server=(local);Data Source=.\SQLEXPRESS;Initial Catalog=db;Integrated Security=true;User ID=sa;Password=123;";

            try
            {
                sqlConn.Open();
            }
            catch
            {
                bConnected = false;
            }
            return bConnected;
        }        
    }    
}
