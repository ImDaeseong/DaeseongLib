using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Xamarin.Forms;
using Xamarin.Forms.Xaml;
using System.Collections.ObjectModel;
using System.Diagnostics;

namespace App1
{
    [XamlCompilation(XamlCompilationOptions.Compile)]
    public partial class AllListView : ContentView
    {
        private GooglePlacesDatabase googlemap = GooglePlacesDatabase.getInstance;
        private ObservableCollection<GooglePlacesItems> mCollection;                
        private List<MapItem> mapAllList;

        public AllListView()
        {
            InitializeComponent();

            mapAllList = null;
            mapAllList = new List<MapItem>();
            MapListview.ItemsSource = mapAllList;
            MapListview.ItemTapped += MapListview_ItemTapped;
                       
            LoadMapData();
        }

        private void IsActIndicator(bool bRun)
        {
            actIndicator.IsRunning = bRun;
            actIndicator.IsVisible = bRun;
        }

        private async void LoadMapData()
        {
            IsActIndicator(true);

            await Task.Factory.StartNew(() =>
            {                
                var mCollection = new ObservableCollection<GooglePlacesItems>(googlemap.GetItems());
                foreach (GooglePlacesItems data in mCollection)
                {
                    mapAllList.Add(new MapItem(data.Name, data.Lat, data.Lng, data.place_id,
                        data.Address, data.phonenumber, data.Internatphonenumber, data.Url, data.Authorname,
                        data.Authorurl, data.Profilephotourl, data.Relativetimedescription, data.TextDesc));
                }
                mCollection.Clear();
            });

            IsActIndicator(false);
        }
  
        private async void MapListview_ItemTapped(object sender, ItemTappedEventArgs e)
        {
            try
            {
                var item = (MapItem)e.Item;
                await Navigation.PushModalAsync(new DetailPage(item));                
            }
            catch
            {
            }
        }

        private async Task<bool> AnimateStackLayout(RoundCornersButton sl)
        {
            await sl.ScaleTo(0.9, 75, Easing.CubicOut);
            await sl.ScaleTo(1, 75, Easing.CubicIn);
            return true;
        }
    }

}