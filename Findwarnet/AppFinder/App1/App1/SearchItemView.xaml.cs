using Xamarin.Forms;
using Xamarin.Forms.Xaml;

namespace App1
{
    [XamlCompilation(XamlCompilationOptions.Compile)]
    public partial class SearchItemView : ContentView
    {
        private string _Name;
        public string Name
        {
            get { return _Name; }
            set { _Name = value; lblName.Text = value; }
        }

        private string _Lat;
        public string Lat
        {
            get { return _Lat; }
            set { _Lat = value; lblLat.Text = value; }
        }

        private string _Lng;
        public string Lng
        {
            get { return _Lng; }
            set { _Lng = value; lblLng.Text = value; }
        }

        private string _place_id;
        public string place_id
        {
            get { return _place_id; }
            set { _place_id = value; lblplace_id.Text = value; }
        }

        private string _Address;
        public string Address
        {
            get { return _Address; }
            set { _Address = value; lblAddress.Text = value; }
        }

        private string _phonenumber;
        public string phonenumber
        {
            get { return _phonenumber; }
            set { _phonenumber = value; lblphonenumber.Text = value; }
        }

        private string _Url;
        public string Url
        {
            get { return _Url; }
            set { _Url = value; lblUrl.Text = value; }
        }

        public SearchItemView(string Name, string Lat, string Lng, string place_id, string Address, string phonenumber, string Url)
        {
            InitializeComponent();

            this.Name = Name;
            this.Lat = Lat;
            this.Lng = Lng;
            this.place_id = place_id;
            this.Address = Address;
            this.phonenumber = phonenumber;
            this.Url = Url;
        }
    }
}