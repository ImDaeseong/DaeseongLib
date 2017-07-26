using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using System.Text.RegularExpressions;


namespace WindowsFormsApplication1
{
    public partial class Form1 : Form
    {
        private bool bComplete;
        private WebBrowser wbSearch;

        public Form1()
        {
            InitializeComponent();

            wbSearch = new WebBrowser();
            wbSearch.AllowWebBrowserDrop = false;
            wbSearch.ScriptErrorsSuppressed = true;
            wbSearch.DocumentCompleted += WbSearch_DocumentCompleted;  
        }
                

        private void WbSearch_DocumentCompleted(object sender, WebBrowserDocumentCompletedEventArgs e)
        {
            if (bComplete) return;

            string sOuterHtml = wbSearch.Document.Body.Parent.OuterHtml;

           
            int nIndex = 0;
            int lastIndex = 0;

            int nTotalIndex = sOuterHtml.Length;
            while (nTotalIndex >= nIndex)
            {
                nIndex = sOuterHtml.IndexOf("<div class=\"thumb\">", lastIndex);
                lastIndex = sOuterHtml.IndexOf("</div>", nIndex);

                string sData = sOuterHtml.Substring(nIndex, (lastIndex - nIndex) );
                if (sData == "") continue;
                //Console.WriteLine(sData);

                //title
                int nTitletKey = sData.IndexOf("alt=") + 4;
                if (nTitletKey == -1) continue;

                int nTitlefirst  = sData.IndexOf("\"", nTitletKey) + 1;
                int nTitleSecond = sData.IndexOf("\"", nTitlefirst);

                string sTitle = sData.Substring(nTitlefirst, (nTitleSecond - nTitlefirst));
                Console.WriteLine(sTitle);
                

                //image
                int nImageKey = sData.IndexOf("src=", nTitletKey) + 4;
                if (nImageKey == -1) continue;

                int nImagefirst  = sData.IndexOf("\"", nImageKey) + 1;
                int nImageSecond = sData.IndexOf("\"", nImagefirst);

                string sImage = sData.Substring(nImagefirst, nImageSecond - nImagefirst);
                Console.WriteLine(sImage);
            }            

        }

        private void InitWebLoad()
        {
            bComplete = false;
            string sUrl = string.Format("http://movie.naver.com/movie/running/current.nhn#");
            wbSearch.Navigate(sUrl);
        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            InitWebLoad();
        }
    }
}
