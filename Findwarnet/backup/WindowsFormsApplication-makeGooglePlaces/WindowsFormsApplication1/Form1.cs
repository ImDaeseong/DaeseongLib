using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Data.SQLite;
using System.IO;
using Newtonsoft.Json;

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
                string sPath = string.Format("C:\\Go\\src\\DaeseongLib\\urlbackup\\urlInfo.txt");
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
                     textBox1.Text += sLog + "\r\n";
                 }
                 )
            );
        }        
    }
}
