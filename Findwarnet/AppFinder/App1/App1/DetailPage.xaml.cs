using System;
using Xamarin.Forms;
using Xamarin.Forms.Xaml;
using System.Threading.Tasks;

namespace App1
{
    [XamlCompilation(XamlCompilationOptions.Compile)]
    public partial class DetailPage : ContentPage
    {
        private MapItem mItem;

        public DetailPage(MapItem item)
        {
            InitializeComponent();

            mItem = item;
            
            NamePcbang.Text = string.Format("Pc방: {0}, 좌표[{1},{2}]", mItem.Name, mItem.Lat, mItem.Lng);
            AddressPcbang.Text = string.Format("주소: {0}",  mItem.Address);

            if (mItem.phonenumber == "" && mItem.phonenumber == "")
                s4.IsVisible = false;
            else
            {
                if (mItem.phonenumber != "")
                    phonenumberPcbang.Text = string.Format("전화번호: {0}", mItem.phonenumber);
                else
                    phonenumberPcbang.Text = string.Format("전화번호: {0}", mItem.Internatphonenumber);
            }

            UrlPcbang.Source = mItem.Url;
        }

        private async void closeDetail_Clicked(object sender, EventArgs e)
        {
            try
            {
                await AnimateStackLayout(closeDetail);
                await Navigation.PopModalAsync();
            }
            catch
            {
            }
        }

        protected override void OnSizeAllocated(double width, double height)
        {
            if (width <= 0 || height <= 0)
                return;

            if (this.Width > this.Height)
            {
                UrlPcbang.WidthRequest = (this.Width - 10);

                if (phonenumberPcbang.Text != "")
                    UrlPcbang.HeightRequest = (this.Height) - (s1.Height + s2.Height + s3.Height + s4.Height + s5.Height);
                else
                    UrlPcbang.HeightRequest = (this.Height) - (s1.Height + s2.Height + s3.Height + s5.Height);
            }
            else
            {
                UrlPcbang.WidthRequest = (this.Width - 10);

                if(phonenumberPcbang.Text != "")
                    UrlPcbang.HeightRequest = (this.Height) - (s1.Height + s2.Height + s3.Height + s4.Height + s5.Height);
                else
                    UrlPcbang.HeightRequest = (this.Height) - (s1.Height + s2.Height + s3.Height + s5.Height);
            }
            
            base.OnSizeAllocated(width, height);
        }

        private async Task<bool> AnimateStackLayout(RoundCornersButton sl)
        {
            await sl.ScaleTo(0.9, 75, Easing.CubicOut);
            await sl.ScaleTo(1, 75, Easing.CubicIn);
            return true;
        }
    }
}