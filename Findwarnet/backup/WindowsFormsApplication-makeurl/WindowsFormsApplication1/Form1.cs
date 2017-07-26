using System;
using System.Data;
using System.Windows.Forms;
using System.IO;
using System.Threading;

namespace WindowsFormsApplication1
{
    public partial class Form1 : Form
    {
        clsDBConnect DBConn = clsDBConnect.getInstance;

        private int nIndex = 0;
        private int nLineCount = 0;
        private StreamWriter wr;
        private static Mutex wmu = new Mutex();

        public Form1()
        {
            InitializeComponent();
            wr = new StreamWriter("data_0.txt", true);
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            if (!DBConn.DbConnection())
            {
                Application.Exit();
                return;
            }
        }

        private void WriteString(string sLine)
        {
            wmu.WaitOne();

            if (nLineCount >= 100)
            {
                nLineCount = 0;
                nIndex++;

                string sFileName = string.Format("data_{0}.txt", nIndex);
                wr = new StreamWriter(sFileName, true);
                wr.WriteLine(sLine);
                wr.Flush();
            }
            else
            {
                nLineCount++;
                wr.WriteLine(sLine);
                wr.Flush();
            }

            wmu.ReleaseMutex();
        }

        private string NullVal(object src, string Value)
        {
            if (src != null)
                return src.ToString();
            return Value;
        }
        
        private void button1_Click(object sender, EventArgs e)
        {
            string sSearch = "";
            string sQuery = "";
            sQuery = "select DISTINCT Query";

            DataSet ds = new DataSet();
            ds = DBConn.GetSelectQuery(sQuery, ds);

            foreach (DataRow dr in ds.Tables[0].Rows)
            {
                string sValue1 = NullVal(dr[0], "").Replace(" ", "").Trim();
                string sValue2 = NullVal(dr[1], "").Replace(" ", "").Trim();
                string sValue3 = NullVal(dr[2], "").Replace(" ", "").Trim();
                string sValue4 = NullVal(dr[3], "").Replace(" ", "").Trim();
                
                if (sValue4 == "")
                    sSearch = string.Format("{0}+{1}+{2}", sValue1, sValue2, sValue3);
                else
                    sSearch = string.Format("{0}+{1}+{2}+{3}", sValue1, sValue2, sValue3, sValue4);

                WriteString(sSearch);
            }
            ds.Dispose();
            wr.Close();

            MessageBox.Show("completed");

            Application.Exit();
        }

        private void Form1_FormClosed(object sender, FormClosedEventArgs e)
        {
            DBConn.DisConnect();
        }
    }
}
