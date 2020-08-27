using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using System.IO;

namespace WindowsFormsApplication1
{
    public partial class Form1 : Form
    {
        bool bComplete = false;

        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            if (textBox1.Text == "") return;

            webBrowser1.Navigate(textBox1.Text);
        }

        private void webBrowser1_DocumentCompleted(object sender, WebBrowserDocumentCompletedEventArgs e)
        {
            //HtmlDocument doc = webBrowser1.Document;
            bComplete = true;
        }

        private void button2_Click(object sender, EventArgs e)
        {
            if (bComplete)
            {
                string shtml = webBrowser1.DocumentText;

                string sHtmlPath = string.Format("{0}\\a.html", Application.StartupPath);
                File.Delete(sHtmlPath);
                
                StreamWriter aStreamWriter = new StreamWriter(sHtmlPath, true);
                aStreamWriter.Write(shtml);
                aStreamWriter.Close();

                MessageBox.Show("저장되었습니다.");
                
                bComplete = false;
            }           

        }

    }
}
