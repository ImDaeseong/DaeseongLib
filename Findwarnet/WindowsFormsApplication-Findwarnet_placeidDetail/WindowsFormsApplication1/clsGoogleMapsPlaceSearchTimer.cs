using System;
using System.Threading;
using Newtonsoft.Json;
using System.Net.Http;
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

        public void crawlerMapurlReverse(object oMethod)
        {
            int nKey = 0;
            bool bResult = false;
            string sUrl = "";
            int nMax = goolgelurl.Count;

            if (nCount == 0)
            {
                nCount = (nMax - 1);
                nKey = nCount;
                sUrl = goolgelurl.Key(nKey).Url;
                bResult = goolgelurl.Key(nKey).Result;
                nCount--;
            }
            else
            {
                nKey = nCount;
                sUrl = goolgelurl.Key(nKey).Url;
                bResult = goolgelurl.Key(nKey).Result;
                nCount--;
            }

            if (!bResult)
                GetPlaceInfo(sUrl, nKey);
        }

        private async void GetPlaceInfo(string url, int nKey)
        {
            try
            {
                string Name = "";
                string Lat = "";
                string Lng = "";
                string Address = "";
                string placeid = "";
                string phonenumber = "";
                string Internatphonenumber = "";
                string Url;
                string Authorname = "";
                string Authorurl = "";
                string Profilephotourl = "";
                string Relativetimedescription = "";
                string TextDesc = "";

                using (var client = new HttpClient())
                {
                    var response = await client.GetStringAsync(url);
                    var result = JsonConvert.DeserializeObject<PlacesApiQueryResponse>(response);

                    if (result.status != "OK")
                    {
                        CFunc.GetMainForm().SetLogText(result.status);
                        return;
                    }

                    Name = string.Format("{0}", result.result.name);
                    Lat = string.Format("{0}", result.result.geometry.location.lat);
                    Lng = string.Format("{0}", result.result.geometry.location.lng);
                    placeid = string.Format("{0}", result.result.place_id);
                    Address = string.Format("{0}", result.result.formatted_address);
                    phonenumber = string.Format("{0}", result.result.formatted_phone_number);
                    Internatphonenumber = string.Format("{0}", result.result.international_phone_number);
                    Url = string.Format("{0}", result.result.url);

                    if (result.result.reviews != null)
                    {
                        foreach (var item in result.result.reviews)
                        {
                            Authorname = string.Format("{0}", item.author_name);
                            Authorurl = string.Format("{0}", item.author_url);
                            Profilephotourl = string.Format("{0}", item.profile_photo_url);
                            Relativetimedescription = string.Format("{0}", item.relative_time_description);
                            TextDesc = string.Format("{0}", item.text);
                        }
                    }
                    
                    lock (mlockWrite)
                    {
                        string sQuery = string.Format("SELECT * FROM tbPlaceId where place_id={0}", CFunc.QStr(placeid));
                        SQLiteDataReader r = DBConn.GetSelectQuery(sQuery);
                        if (!r.Read())
                        {
                            sQuery = "";
                            sQuery = string.Format(
                                @"INSERT INTO tbPlaceId ( Name, Lat, Lng, place_id, Address, phonenumber,
                                  Internatphonenumber, Url, Authorname, Authorurl, Profilephotourl, 
                                  Relativetimedescription, TextDesc)
                                  VALUES ({0},{1},{2},{3},{4},{5},{6},{7},{8},{9},{10},{11},{12});", 
                                CFunc.QStr(Name), CFunc.QStr(Lat), CFunc.QStr(Lng), CFunc.QStr(placeid),
                                CFunc.QStr(Address), CFunc.QStr(phonenumber),CFunc.QStr(Internatphonenumber),
                                CFunc.QStr(Url), CFunc.QStr(Authorname), CFunc.QStr(Authorurl),
                                CFunc.QStr(Profilephotourl),CFunc.QStr(Relativetimedescription),CFunc.QStr(TextDesc));
                            DBConn.ExeCuteQuery(sQuery);
                            CFunc.GetMainForm().SetLogText("Insert:" + sQuery);//Console.WriteLine("Insert:" + sQuery);
                            goolgelurl.Key(nKey).Result = true;
                        }
                        else
                        {
                            sQuery = "";
                            sQuery = string.Format(
                                @"UPDATE tbPlaceId SET Name ={0}, Lat ={1}, Lng ={2}, Address ={3}, phonenumber ={4},
                                  Internatphonenumber ={5}, Url ={6}, Authorname ={7}, Authorurl ={8}, Profilephotourl ={9}, 
                                  Relativetimedescription ={10}, TextDesc ={11} where place_id={12}",
                                 CFunc.QStr(Name), CFunc.QStr(Lat), CFunc.QStr(Lng), CFunc.QStr(Address), CFunc.QStr(phonenumber),
                                 CFunc.QStr(Internatphonenumber), CFunc.QStr(Url), CFunc.QStr(Authorname), CFunc.QStr(Authorurl),
                                 CFunc.QStr(Profilephotourl), CFunc.QStr(Relativetimedescription), CFunc.QStr(TextDesc), CFunc.QStr(placeid));
                            DBConn.ExeCuteQuery(sQuery);
                            CFunc.GetMainForm().SetLogText("Update:" + sQuery);
                            goolgelurl.Key(nKey).Result = true;
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
            //tCallback = new TimerCallback(this.crawlerMapurlReverse);
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

    public class reviews
    {
        public string author_name { get; set; }
        public string author_url { get; set; }
        public string profile_photo_url { get; set; }
        public string relative_time_description { get; set; }
        public string text { get; set; }
    }

    public class Result
    {
        public string formatted_address { get; set; }
        public string formatted_phone_number { get; set; }        
        public Geometry geometry { get; set; }
        public string icon { get; set; }
        public string id { get; set; }
        public string international_phone_number { get; set; }
        public string name { get; set; }       
        public string place_id { get; set; }
        public string url { get; set; }        
        public List<reviews> reviews { get; set; }        
    }

    public class PlacesApiQueryResponse
    {
        public List<object> html_attributions { get; set; }
        public Result result { get; set; }
        public string status { get; set; }
    }
}
