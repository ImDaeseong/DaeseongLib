using System;
using System.Collections.Generic;
using System.Windows.Forms;
using Newtonsoft.Json;
using System.Net.Http;
using System.IO;
using System.Threading;

namespace WindowsFormsApplication1
{
    public partial class Form1 : Form
    {
        private StreamWriter wr;
        private static Mutex wmu = new Mutex();

        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            wr = new StreamWriter("mapInfo.txt", true);
        }

        private void WriteString(string sLine)
        {
            wmu.WaitOne();
            wr.WriteLine(sLine);
            wr.Flush();
            wmu.ReleaseMutex();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            string sInput = "search workd";
            string sKey = "API Key";
            string sWord = "kind";           
            string sSearch = "";
            sSearch = sInput.Trim();
            sSearch = sInput.Replace(" ", "+");
            string url = string.Format("https://maps.googleapis.com/maps/api/place/textsearch/json?query={0}+{1}&key={2}", sWord, sSearch, sKey);
            //Console.WriteLine(url);

            GetPlaceInfo(url);
            MessageBox.Show("Completed");
        }

        private async void GetPlaceInfo(string url)
        {
            using (var client = new HttpClient())
            {
                var response = await client.GetStringAsync(url);
                var result = JsonConvert.DeserializeObject<PlacesApiQueryResponse>(response);
                //Console.WriteLine(result.status);
                
                if (result.status != "OK") return;
                
                foreach (var i in result.results)
                {
                    string sMap = string.Format("{0}|{1}|{2}", i.name, i.geometry.location.lat, i.geometry.location.lng);
                    WriteString(sMap);
                }
            }
        }

    }


    public class Location
    {
        public double lat { get; set; }
        public double lng { get; set; }
    }

    public class Geometry
    {
        public Location location { get; set; }
    }

    public class OpeningHours
    {
        public bool open_now { get; set; }
        public List<object> weekday_text { get; set; }
    }

    public class Photo
    {
        public int height { get; set; }
        public List<string> html_attributions { get; set; }
        public string photo_reference { get; set; }
        public int width { get; set; }
    }

    public class Result
    {
        public Geometry geometry { get; set; }
        public string icon { get; set; }
        public string id { get; set; }
        public string name { get; set; }
        public OpeningHours opening_hours { get; set; }
        public List<Photo> photos { get; set; }
        public string place_id { get; set; }
        public double rating { get; set; }
        public string reference { get; set; }
        public string scope { get; set; }
        public List<string> types { get; set; }
        public string vicinity { get; set; }
    }

    public class PlacesApiQueryResponse
    {
        public List<object> html_attributions { get; set; }
        public List<Result> results { get; set; }
        public string status { get; set; }
    }

}
