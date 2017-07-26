using System;
using System.Collections.Generic;
using System.Windows.Forms;
using Newtonsoft.Json;
using System.Net.Http;

/*
https://www.bithumb.com/u1/US127
*/

namespace WindowsFormsApplication1
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private async void getbithumbTicker(string url)
        {
            using (var client = new HttpClient())
            {
                var resp = await client.GetStringAsync(url);
                var result = JsonConvert.DeserializeObject<Response>(resp);
                //Console.WriteLine(result.Status);

                if (result.Status != "0000") return;

                Console.WriteLine(result.Data.Opening_price);
                Console.WriteLine(result.Data.Closing_price);
                Console.WriteLine(result.Data.Min_price);
                Console.WriteLine(result.Data.Max_price);
                Console.WriteLine(result.Data.Average_price);
                Console.WriteLine(result.Data.Units_traded);
                Console.WriteLine(result.Data.Volume_1day);
                Console.WriteLine(result.Data.Volume_7day);
                Console.WriteLine(result.Data.Buy_price);
                Console.WriteLine(result.Data.Sell_price);
                Console.WriteLine(result.Data.Date);
            }
        }

        private void button1_Click(object sender, EventArgs e)
        {
            string sUrl = "https://api.bithumb.com/public/ticker/BTC";
            getbithumbTicker(sUrl);
        }


        private async void getbithumborderbook(string url)
        {
            using (var client = new HttpClient())
            {
                var resp = await client.GetStringAsync(url);
                var result = JsonConvert.DeserializeObject<Response1>(resp);
                //Console.WriteLine(result.Status);

                if (result.Status != "0000") return;

                Console.WriteLine(result.Data.Timestamp);
                Console.WriteLine(result.Data.Order_currency);
                Console.WriteLine(result.Data.Payment_currency);

                foreach (var val in result.Data.Bids)
                {
                    Console.WriteLine(val.Quantity);
                    Console.WriteLine(val.Price);
                }

                foreach (var val in result.Data.Asks)
                {
                    Console.WriteLine(val.Quantity);
                    Console.WriteLine(val.Price);
                }
            }
        }

        private void button2_Click(object sender, EventArgs e)
        {
            string sUrl = "https://api.bithumb.com/public/orderbook/BTC";
            getbithumborderbook(sUrl);
        }

        private async void getbithumbrecent_transactions(string url)
        {
            using (var client = new HttpClient())
            {
                var resp = await client.GetStringAsync(url);
                var result = JsonConvert.DeserializeObject<Response2>(resp);
                //Console.WriteLine(result.Status);

                if (result.Status != "0000") return;

                foreach (var val in result.Data)
                {
                    Console.WriteLine(val.Transaction_date);
                    Console.WriteLine(val.Type);
                    Console.WriteLine(val.Units_traded);
                    Console.WriteLine(val.Price);
                    Console.WriteLine(val.Total);
                }
            }
        }

        private void button3_Click(object sender, EventArgs e)
        {
            string sUrl = "https://api.bithumb.com/public/recent_transactions/BTC";
            getbithumbrecent_transactions(sUrl);
        }
    }


    public class result_data
    {
        public string Opening_price { get; set; }
        public string Closing_price { get; set; }
        public string Min_price { get; set; }
        public string Max_price { get; set; }
        public string Average_price { get; set; }
        public string Units_traded { get; set; }
        public string Volume_1day { get; set; }
        public string Volume_7day { get; set; }
        public string Buy_price { get; set; }
        public string Sell_price { get; set; }
        public string Date { get; set; }
    }
    public class Response
    {
        public string Status { get; set; }
        public result_data Data { get; set; }
    }



    public class price_data
    {
        public string Quantity { get; set; }
        public string Price { get; set; }
    }
    public class result_data1
    {
        public string Timestamp { get; set; }

        public string Order_currency { get; set; }

        public string Payment_currency { get; set; }

        public List<price_data> Bids { get; set; }

        public List<price_data> Asks { get; set; }
    }
    public class Response1
    {
        public string Status { get; set; }
        public result_data1 Data { get; set; }
    }
    


    public class transaction_data
    {
        public string Transaction_date { get; set; }
        public string Type { get; set; }

        public string Units_traded { get; set; }

        public string Price { get; set; }

        public string Total { get; set; }
    }
    public class Response2
    {
        public string Status { get; set; }
        public List<transaction_data> Data { get; set; }
    }

}
