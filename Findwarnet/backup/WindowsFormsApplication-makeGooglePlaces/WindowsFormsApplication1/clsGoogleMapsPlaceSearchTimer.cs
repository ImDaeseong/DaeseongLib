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
            string Name;
            string Lat;
            string Lng;
            string Address;
            string placeid;

            using (var client = new HttpClient())
            {
                var response = await client.GetStringAsync(url);
                var result = JsonConvert.DeserializeObject<PlacesApiQueryResponse>(response);
                //Console.WriteLine(result.status);

                if (result.status != "OK")
                {
                    CFunc.GetMainForm().SetLogText(result.status);
                    return;
                }

                foreach (var i in result.results)
                {
                    Name = string.Format("{0}", i.name);
                    Lat = string.Format("{0}", i.geometry.location.lat);
                    Lng = string.Format("{0}", i.geometry.location.lng);
                    placeid = string.Format("{0}", i.place_id);
                    Address = string.Format("{0}", i.formatted_address);
                   
                    lock (mlockWrite)
                    {
                        string sQuery = string.Format("SELECT * FROM FormattedAddress where Lat={0} and Lng={1}", CFunc.QStr(Lat), CFunc.QStr(Lng));
                        SQLiteDataReader r = DBConn.GetSelectQuery(sQuery);
                        if (!r.Read())
                        {
                            sQuery = "";
                            sQuery = string.Format("INSERT INTO FormattedAddress ( Name, Lat, Lng, place_id, Address ) VALUES ({0},{1},{2},{3},{4});", CFunc.QStr(Name), CFunc.QStr(Lat), CFunc.QStr(Lng), CFunc.QStr(placeid), CFunc.QStr(Address));
                            DBConn.ExeCuteQuery(sQuery);
                            CFunc.GetMainForm().SetLogText("Insert:" + sQuery);//Console.WriteLine("Insert:" + sQuery);
                            goolgelurl.Key(nKey).Result = true;
                        }
                    }
                }
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
