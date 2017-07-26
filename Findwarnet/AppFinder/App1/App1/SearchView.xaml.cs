using System;
using System.Collections.ObjectModel;
using Xamarin.Forms;
using Xamarin.Forms.Maps;
using Xamarin.Forms.Xaml;

namespace App1
{
    [XamlCompilation(XamlCompilationOptions.Compile)]
    public partial class SearchView : ContentView
    {
        private GooglePlacesDatabase googlemap = GooglePlacesDatabase.getInstance;
        private ObservableCollection<GooglePlacesItems> mCollection;

        public SearchView()
        {
            InitializeComponent();
        }

        private void eSearch_Completed(object sender, EventArgs e)
        {
            try
            {
                var eSearch = (Entry)sender;
                string sSearch = eSearch.Text;

                if (sSearch == "") return;

                var ItemTapped = new TapGestureRecognizer();
                ItemTapped.Tapped += ItemTapped_Tapped;

                stackList.Children.Clear();

                mCollection = new ObservableCollection<GooglePlacesItems>(googlemap.GetSearchAddress(sSearch));

                foreach (GooglePlacesItems data in mCollection)
                {
                    var item = new SearchItemView(data.Name, data.Lat, data.Lng, data.place_id, data.Address, data.phonenumber, data.Url);
                    item.GestureRecognizers.Add(ItemTapped);
                    stackList.Children.Add(item);
                }
                mCollection.Clear();
                eSearch.Text = "";

                this.UpdateChildrenLayout();
                this.InvalidateLayout();
            }
            catch
            {
            }
        }

        private void ItemTapped_Tapped(object sender, EventArgs e)
        {
            try
            {
                map.Pins.Clear();

                var item = (SearchItemView)sender;
                string sPin = string.Format("[{0}]{1}", item.Name, item.Address);

                var location = new Position(double.Parse(item.Lat), double.Parse(item.Lng));
                map.MoveToRegion(MapSpan.FromCenterAndRadius(location, Distance.FromKilometers(0.3)));
                var pin = new Pin
                {
                    Type = PinType.Place,
                    Position = location,
                    Label = sPin
                };
                map.Pins.Add(pin);
                SetFrameVisible(false);

                this.UpdateChildrenLayout();
                this.InvalidateLayout();
            }
            catch
            {
            }
        }

        private void SetFrameVisible(bool bVisible)
        {
            frame1.IsVisible = bVisible;
            frame2.IsVisible = !bVisible;
        }

        protected override void LayoutChildren(double x, double y, double width, double height)
        {
            base.LayoutChildren(x, y, width, height);
        }

        protected override void OnSizeAllocated(double width, double height)
        {
            if (map.IsVisible)
            {
                map.WidthRequest = (this.Width - 10);
                map.HeightRequest = (this.Height - eSearch.Height - 10);
            }

            base.OnSizeAllocated(width, height);
        }

        private void eSearch_TextChanged(object sender, TextChangedEventArgs e)
        {
            try
            {
                SetFrameVisible(true);
            }
            catch
            {
            }
        }
    }
}