using System;

namespace App1
{
    public class MapItem
    {
        public string Name { get; set; }
        public string Lat { get; set; }
        public string Lng { get; set; }
        public string place_id { get; set; }
        public string Address { get; set; }
        public string phonenumber { get; set; }
        public string Internatphonenumber { get; set; }
        public string Url { get; set; }
        public string Authorname { get; set; }
        public string Authorurl { get; set; }
        public string Profilephotourl { get; set; }
        public string Relativetimedescription { get; set; }
        public string TextDesc { get; set; }

        public string DisplayName { get; set; }

        public MapItem(string Name, string Lat, string Lng, string place_id,
            string Address, string phonenumber, string Internatphonenumber,
            string Url, string Authorname, string Authorurl, string Profilephotourl,
            string Relativetimedescription, string TextDesc)
        {
            this.Name = Name;
            this.Lat = Lat;
            this.Lng = Lng;
            this.place_id = place_id;
            this.Address = Address;
            this.phonenumber = phonenumber;
            this.Internatphonenumber = Internatphonenumber;
            this.Url = Url;
            this.Authorname = Authorname;
            this.Authorurl = Authorurl;
            this.Profilephotourl = Profilephotourl;
            this.Relativetimedescription = Relativetimedescription;
            this.TextDesc = TextDesc;
            this.DisplayName = string.Format("[{0}] {1}", Name, Address);
        }
    }
}
