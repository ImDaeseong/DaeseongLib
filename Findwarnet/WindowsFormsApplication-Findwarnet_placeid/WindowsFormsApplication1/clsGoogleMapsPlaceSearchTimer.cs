using System;
using System.Threading;
using System.Windows.Forms;
using System.Diagnostics;
using System.Text;
using Newtonsoft.Json;
using System.Net.Http;
using System.IO;
using Timer = System.Threading.Timer;
using System.Collections.Generic;
using System.Data.SQLite;

namespace WindowsFormsApplication1
{
    class clsGoogleMapsPlaceSearchTimer
    {
        clsDBConnect DBConn = clsDBConnect.getInstance;
        clsUrlList goolgelurl = clsUrlList.getInstance;
        clsComm CFunc = clsComm.getInstance;

        TimerCallback tCallback;
        Timer SearchMapTimer;
        private int nCount = 0;

        static readonly object mlockWrite = new object();

        public void crawlerMapurl(object oMethod)
        {
            int nKey = 0;
            bool bResult = false;
            string sUrl = "";
            int nMax = goolgelurl.Count;
            if (nCount <= (nMax -1))
            {
                nKey = nCount;
                sUrl = goolgelurl.Key(nKey).Url;
                bResult = goolgelurl.Key(nKey).Result;
                nCount++;
            }
            else
            {
                nCount = 0;
                nKey = nCount;
                sUrl = goolgelurl.Key(nKey).Url;
                bResult = goolgelurl.Key(nKey).Result;
                nCount++;
            }

            if (!bResult)
                GetPlaceInfo(sUrl, nKey);
        }

        private async void GetPlaceInfo(string url, int nKey)
        {
            try
            {

                string placeid;

                using (var client = new HttpClient())
                {
                    var response = await client.GetStringAsync(url);
                    var result = JsonConvert.DeserializeObject<PlacesApiQueryResponse>(response);

                    if (result.status != "OK")
                    {
                        CFunc.GetMainForm().SetLogText(result.status);
                        return;
                    }

                    foreach (var i in result.results)
                    {
                        placeid = string.Format("{0}", i.place_id);

                        lock (mlockWrite)
                        {
                            string sQuery = string.Format("SELECT * FROM tbPlaceId where place_id={0}", CFunc.QStr(placeid));
                            SQLiteDataReader r = DBConn.GetSelectQuery(sQuery);
                            if (!r.Read())
                            {
                                sQuery = "";
                                sQuery = string.Format("INSERT INTO tbPlaceId (place_id) VALUES ({0});", CFunc.QStr(placeid));
                                DBConn.ExeCuteQuery(sQuery);
                                CFunc.GetMainForm().SetLogText("Insert:" + sQuery);
                                goolgelurl.Key(nKey).Result = true;
                            }
                        }
                    }
                }
            }
            catch (Exception ex)
            {
                CFunc.GetMainForm().SetLogText("Insert Error:" + ex.Message.ToString());
            }
        }
        
        public void Start()
        {
            tCallback = new TimerCallback(this.crawlerMapurl);
            SearchMapTimer = new Timer(tCallback, this, 0, 10000);//10초마다 
        }

        public void Stop()
        {
            if (SearchMapTimer != null)
            {
                SearchMapTimer.Dispose();
            }
        }

        public clsGoogleMapsPlaceSearchTimer()
        {
        }

        ~ clsGoogleMapsPlaceSearchTimer()
        {
            Stop();
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
        public string formatted_address { get; set; }
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
