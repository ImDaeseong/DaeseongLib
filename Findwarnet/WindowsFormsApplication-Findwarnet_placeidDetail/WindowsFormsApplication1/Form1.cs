using System;
using System.Windows.Forms;
using System.IO;

namespace WindowsFormsApplication1
{
    public partial class Form1 : Form
    {
        clsGoogleMapsPlaceSearchTimer tMapsTimer = null;

        clsUrlList goolgelurl = clsUrlList.getInstance;
        clsDBConnect DBConn = clsDBConnect.getInstance;
        clsComm CFunc = clsComm.getInstance;

        static readonly object mlockWrite = new object();
        
        private bool Loadformatted_address()
        {            
            bool bLoad = true;
            try
            {
                int nIndex = 0;
                string sUrl = "";
                string sPlaceID = "";
                string sApikey = "Api key";
                string sPath = string.Format("GooglePlacesInfo.txt");
                StreamReader reader = new StreamReader(sPath);
                using (reader)
                {
                    string line = reader.ReadLine();
                    while (line != null)
                    {                            
                        lock (mlockWrite)
                        {
                            sPlaceID = CFunc.NullVal(line, "").Replace("\"", "").Replace("\r", "").Replace("\n", "").Replace("\t", "").Trim();
                            sUrl = string.Format("https://maps.googleapis.com/maps/api/place/details/json?placeid={0}&key={1}&language=ko", sPlaceID, sApikey);
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
            tMapsTimer.Stop();
            Close();
        }

        public void SetLogText(string sLog)
        {
            this.Invoke(new MethodInvoker(
                 delegate ()
                 {
                     textBox1.AppendText(sLog + "\r\n");
                 }
                 )
            );
        }        
    }
}
