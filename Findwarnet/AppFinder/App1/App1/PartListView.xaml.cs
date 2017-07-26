using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Xamarin.Forms;
using Xamarin.Forms.Xaml;
using System.Collections.ObjectModel;

namespace App1
{
    [XamlCompilation(XamlCompilationOptions.Compile)]
    public partial class PartListView : ContentView
    {
        private GooglePlacesDatabase googlemap = GooglePlacesDatabase.getInstance;
        private ObservableCollection<GooglePlacesItems> mCollection;

        private const int ListCount = 20;
        private int nTotalCount;
        private int total_page;
        private int nCurrentPage;

        private List<MapItem> mapAllList;

        public PartListView()
        {
            InitializeComponent();

            mapAllList = null;
            mapAllList = new List<MapItem>();
            MapListview.ItemsSource = mapAllList;
            MapListview.ItemTapped += MapListview_ItemTapped;

            MapListview.IsPullToRefreshEnabled = true;

            LoadMapDataPart();
        }

        private void IsActIndicator(bool bRun)
        {
            actIndicator.IsRunning = bRun;
            actIndicator.IsVisible = bRun;
        }

        private async void LoadMapDataPart()
        {
            IsActIndicator(true);

            await Task.Factory.StartNew(() =>
            {
                mCollection = new ObservableCollection<GooglePlacesItems>(googlemap.GetItems());

                nTotalCount = mCollection.Count;
                total_page = ((nTotalCount - 1) / ListCount + 1);

                nCurrentPage = 0;
                SelectPage(nCurrentPage, true);
            });

            IsActIndicator(false);
        }

        private void SelectPage(int nCurrentPage, bool bFirst = false)
        {
            try
            {
                int nStartIndex = ListCount * nCurrentPage;

                int nListCount = 0;
                if (nCurrentPage == (total_page - 1))
                    nListCount = nTotalCount - (ListCount * nCurrentPage);
                else
                    nListCount = ListCount;

                mapAllList.Clear();
                if (!bFirst) MapListview.ItemsSource = null;
                MapListview.ItemsSource = mapAllList;

                for (int k = 0; k < nListCount; k++)
                {
                    mapAllList.Add(new MapItem(mCollection[nStartIndex].Name,
                        mCollection[nStartIndex].Lat,
                        mCollection[nStartIndex].Lng,
                        mCollection[nStartIndex].place_id,
                        mCollection[nStartIndex].Address,
                        mCollection[nStartIndex].phonenumber,
                        mCollection[nStartIndex].Internatphonenumber,
                        mCollection[nStartIndex].Url,
                        mCollection[nStartIndex].Authorname,
                        mCollection[nStartIndex].Authorurl,
                        mCollection[nStartIndex].Profilephotourl,
                        mCollection[nStartIndex].Relativetimedescription,
                        mCollection[nStartIndex].TextDesc));
                    nStartIndex++;
                }
            }
            catch
            {
            }
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

        private async void PreMap_Clicked(object sender, EventArgs e)
        {
            try
            {
                nCurrentPage = nCurrentPage - 1;
                if (nCurrentPage < 0)
                    nCurrentPage = 0;

                SelectPage(nCurrentPage);

                await MovePrePage();
                await BannerStackLayout(true);
                await AnimateStackLayout(PreMap);
            }
            catch
            {
            }
        }

        private async void nexMap_Clicked(object sender, EventArgs e)
        {
            try
            {
                nCurrentPage = nCurrentPage + 1;
                if (total_page <= nCurrentPage)
                    nCurrentPage = total_page - 1;

                SelectPage(nCurrentPage);

                await MoveNextPage();
                await BannerStackLayout(false);
                await AnimateStackLayout(nexMap);
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

        private async Task<bool> BannerStackLayout(bool bLeft)
        {
            if (bLeft)
            {
                await bannerImg.LayoutTo(new Rectangle(0, 0, 100, 40), 0);
                bannerImg.IsVisible = true;
                await bannerImg.LayoutTo(new Rectangle() { X = 1000, Y = 0, Width = 100, Height = 40 }, 4000);
            }
            else
            {
                await bannerImg.LayoutTo(new Rectangle(Width - 100, 0, 100, 40), 0);
                bannerImg.IsVisible = true;
                await bannerImg.LayoutTo(new Rectangle() { X = -100, Y = 0, Width = 100, Height = 40 }, 2000);
            }
            return true;
        }


        private async Task<bool> MovePrePage()
        {
            await MapListview.TranslateTo(-1000, 0);
            await MapListview.FadeTo(0, 1);
            await MapListview.TranslateTo(1000, 0);
            await MapListview.FadeTo(1, 1);
            await MapListview.TranslateTo(0, 0);
            return true;
        }

        private async Task<bool> MoveNextPage()
        {
            await MapListview.TranslateTo(1000, 0);
            await MapListview.FadeTo(0, 1);
            await MapListview.TranslateTo(-1000, 0);
            await MapListview.FadeTo(1, 1);
            await MapListview.TranslateTo(0, 0);
            return true;
        }

    }
}