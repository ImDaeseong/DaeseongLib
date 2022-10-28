using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using Newtonsoft.Json;
using System.Net.Http;
using static System.Windows.Forms.VisualStyles.VisualStyleElement;
using System.Security.Policy;
using System.Xml;
using static System.Resources.ResXFileRef;
using System.Xml.Linq;
using Newtonsoft.Json.Linq;
using static System.Windows.Forms.AxHost;

namespace WindowsFormsApp1
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            textBox1.Text = "";

            string sUrl = GetUrlString(1);
            loadUrl_Xml(sUrl);
        }

        private void button2_Click(object sender, EventArgs e)
        {
            textBox1.Text = "";

            string sUrl = GetUrlString(2);
            loadUrl_JSON(sUrl);
        }

        private string GetUrlString(int nType)
        {
            string spage = "https://apis.data.go.kr/1360000/AsosHourlyInfoService/getWthrDataList?";

            string serviceKey = "%2FSWbuoncrZtSM3DaBUA4PJVxqJMFKs0Eu%2F%2FzgFQf8dvVjzIi8ESOjmRaQtAkLKoQUS3S%2BZy%2FwLwR08%2BCT9BWuA%3D%3D";

            int pageNo = 1;

            int numOfRows = 10;


            string dataType = "";    
            if (nType == 1) {
                dataType = "XML";    
            }
            else
            {
                dataType = "JSON";    
            }

            string dataCd = "ASOS";

            string dateCd = "HR";

            //string sDay = string.Format("{0:yyyyMMdd}", DateTime.Now.AddDays(-1));
            //string sMonth = string.Format("{0:yyyyMMdd}", DateTime.Now.AddMonths(-1));

            string startDt = string.Format("{0:yyyyMMdd}", DateTime.Now.AddDays(-1)); // 하루전

            string startHh = "01";

            string endDt = string.Format("{0:yyyyMMdd}", DateTime.Now.AddDays(-1)); // 하루전

            string endHh = "01";

            string stnIds = "108"; // 서울

            string sUrl = string.Format("{0}serviceKey={1}&pageNo={2}&numOfRows={3}&dataType={4}&dataCd={5}&dateCd={6}&startDt={7}&startHh={8}&endDt={9}&endHh={10}&stnIds={11}", spage, serviceKey, pageNo, numOfRows, dataType, dataCd, dateCd, startDt, startHh, endDt, endHh, stnIds);
            return sUrl;
        }

        private async void loadUrl_Xml(string url)
        {
            using (var client = new HttpClient())
            {
                string response = await client.GetStringAsync(url);

                XmlDocument xml = new XmlDocument();
                xml.LoadXml(response);

                XmlNode xNode = xml.SelectSingleNode("response/body/items/item");
                if(xNode != null)
                {
                    foreach (XmlNode item in xNode.ChildNodes)
                    {
                        //Console.WriteLine(item.Name + " :" + item.InnerText);

                        if (item.Name == "tm")
                        {
                            textBox1.Text += string.Format("시간:{0}", item.InnerText) + "\r\n";
                        }
                        else if (item.Name == "stnId")
                        {
                            textBox1.Text += string.Format("지점번호:{0}", item.InnerText) + "\r\n";
                        }
                        else if (item.Name == "stnNm")
                        {
                            textBox1.Text += string.Format("지점명:{0}", item.InnerText) + "\r\n";
                        }
                        else if (item.Name == "ta")
                        {
                            textBox1.Text += string.Format("기온:{0}", item.InnerText) + "\r\n";
                        }
                        else if (item.Name == "ws")
                        {
                            textBox1.Text += string.Format("풍속:{0}", item.InnerText) + "\r\n";
                        }
                        else if (item.Name == "rn")
                        {
                            textBox1.Text += string.Format("강수량:{0}", item.InnerText) + "\r\n";
                        }
                        else if (item.Name == "wd")
                        {
                            textBox1.Text += string.Format("풍향:{0}", item.InnerText) + "\r\n";
                        }
                        else if (item.Name == "hm")
                        {
                            textBox1.Text += string.Format("습도:{0}", item.InnerText) + "\r\n";
                        }
                    }
                }
            }
        }

        private async void loadUrl_JSON(string url)
        {
            using (var client = new HttpClient())
            {
                string response = await client.GetStringAsync(url);
                //Console.WriteLine(response);

                JObject jObject = JObject.Parse(response);
                if (jObject == null) return;

                JToken jToken = jObject["response"]["body"]["items"]["item"];
                if (jToken == null) return;

                foreach (JToken item in jToken)
                {
                    //Console.WriteLine(item.ToString());

                    textBox1.Text += string.Format("시간:{0}", item["tm"]) + "\r\n";

                    textBox1.Text += string.Format("지점번호:{0}", item["stnId"]) + "\r\n";

                    textBox1.Text += string.Format("지점명:{0}", item["stnNm"]) + "\r\n";

                    textBox1.Text += string.Format("기온:{0}", item["ta"]) + "\r\n";

                    textBox1.Text += string.Format("풍속:{0}", item["ws"]) + "\r\n";

                    textBox1.Text += string.Format("강수량:{0}", item["rn"]) + "\r\n";

                    textBox1.Text += string.Format("풍향:{0}", item["wd"]) + "\r\n";

                    textBox1.Text += string.Format("습도:{0}", item["hm"]) + "\r\n";
                }

            }
        }
    }
}
