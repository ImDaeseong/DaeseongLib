using SQLite;

namespace App1
{
    [Table("tbPlaceId")]
    public class GooglePlacesItems
    {
        public GooglePlacesItems()
        {
        }

        [PrimaryKey, AutoIncrement]
        public int rIndex { get; set; }

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
    }
}
