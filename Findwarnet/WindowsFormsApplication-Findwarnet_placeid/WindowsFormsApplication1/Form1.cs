using System;
using System.Windows.Forms;
using System.Data.SQLite;
using System.IO;
using System.Threading;

namespace WindowsFormsApplication1
{
    public partial class Form1 : Form
    {
        clsGoogleMapsPlaceSearchTimer tMapsTimer = null;

        clsUrlList goolgelurl = clsUrlList.getInstance;
        clsDBConnect DBConn = clsDBConnect.getInstance;
        clsComm CFunc = clsComm.getInstance;

        static readonly object mlockWrite = new object();

        private StreamWriter wr;
        private static Mutex wmu = new Mutex();

        private bool Loadformatted_address()
        {            
            bool bLoad = true;
            try
            {
                int nIndex = 0;
                string sUrl = "";
                string sPath = string.Format("urlInfo.txt");
                StreamReader reader = new StreamReader(sPath);
                using (reader)
                {
                    string line = reader.ReadLine();
                    while (line != null)
                    {                            
                        lock (mlockWrite)
                        {
                            sUrl = CFunc.NullVal(line, "").Replace("\"", "").Replace("\r", "").Replace("\n", "").Replace("\t", "").Trim();
                            sUrl += "&language=ko";
                            goolgelurl.Add(nIndex, sUrl);
                            nIndex++;
                        }
                        line = reader.ReadLine();
                    }
                }
            }
            catch (Exception ex)
            {
                bLoad = false;
                Console.WriteLine(ex.Message.ToString());
            }
            return bLoad;
        }
        
        public Form1()
        {
            InitializeComponent();
        }
        
        private void Form1_Load(object sender, EventArgs e)
        {
            if (DBConn.DbConnection())
            {
                //DBConn.createTableGooglePlaces();
                Console.WriteLine("connected");
            }
            else
            {
                Console.WriteLine("not connected");                
            }

            //DBConn.ReadDataFile();
        }
        
        private void button1_Click(object sender, EventArgs e)
        {
            if (Loadformatted_address())
            {
                Console.WriteLine("Load Count:" + goolgelurl.Count);
                Console.WriteLine("Start Timer");

                tMapsTimer = new clsGoogleMapsPlaceSearchTimer();
                tMapsTimer.Start();
            }
        }

        private void button2_Click(object sender, EventArgs e)
        {
            if(tMapsTimer != null)
            {
                tMapsTimer.Stop();
            }
            
            Close();
        }

        public void SetLogText(string sLog)
        {
            this.Invoke(new MethodInvoker(
                 delegate ()
                 {
                     textBox1.AppendText(sLog + "\r\n");
                     //textBox1.Text += sLog + "\r\n";
                 }
                 )
            );
        }

        private void WriteString(string sLine)
        {
            wmu.WaitOne();
            wr.WriteLine(sLine);
            wr.Flush();
            wmu.ReleaseMutex();
        }

        private void button3_Click(object sender, EventArgs e)
        {
            try
            {
                int rowCount = 0;
                wr = new StreamWriter("GooglePlacesInfo.txt", true);
                string sQuery = string.Format("SELECT * FROM tbPlaceId");
                SQLiteDataReader r = DBConn.GetSelectQuery(sQuery);
                while (r.Read())
                {
                    rowCount++;
                    string place_id = string.Format("{0}", r["place_id"].ToString());
                    WriteString(place_id);
                }
                wr.Close();

                string sMeg = string.Format("{0} 개 생성", rowCount.ToString());
                SetLogText(sMeg);
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());
            }
        }
    }
}
