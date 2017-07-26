using System;
using System.Collections;

namespace WindowsFormsApplication1
{
    class clsUrlList : DictionaryBase
    {
        private static clsUrlList selfInstance = null;
        public static clsUrlList getInstance
        {
            get
            {
                if (selfInstance == null) selfInstance = new clsUrlList();
                return selfInstance;
            }
        }

        public clsUrlList()
        {
            Dictionary.Clear();
        }

        ~clsUrlList()
        {
            Dictionary.Clear();
            this.Clear();

        }

        public urlTable Key(int nKey)
        {
            return (urlTable)Dictionary[nKey];
        }
        
        public void Remove(int nKey)
        {
            Dictionary.Remove(nKey);
        }

        public void Add(int nIndex, string sUrl, bool bResult = false)
        {
            urlTable dTable = new urlTable();
            dTable.RIndex = nIndex;
            dTable.Url = sUrl;
            dTable.Result = bResult;

            try
            {
                Dictionary.Add(nIndex, dTable);
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message.ToString());
            }
        }
    }

    class urlTable
    {
        private int m_rIndex = 0;
        private string m_sUrl = "";
        private bool m_bResult = false;

        public urlTable()
        {
            Clear();
        }

        public void Clear()
        {
            m_rIndex = 0;
            m_sUrl = "";
        }

        public int RIndex
        {
            get { return m_rIndex; }
            set { m_rIndex = value; }
        }

        public string Url
        {
            get { return m_sUrl; }
            set { m_sUrl = value; }
        }

        public bool Result
        {
            get { return m_bResult; }
            set { m_bResult = value; }            
        }
    }
}
