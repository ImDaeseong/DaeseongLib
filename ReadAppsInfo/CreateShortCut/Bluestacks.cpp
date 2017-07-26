#include "stdafx.h"
#include "Bluestacks.h"

Bluestacks::Bluestacks()
{
	DeleteAllMemory();
}

Bluestacks::~Bluestacks()
{
	DeleteAllMemory();
}

CString Bluestacks::GetModulePath()
{
	char cTemp[MAX_PATH];
	char *spzRealDirEnd;
	CString strModulePath;

	GetModuleFileName(NULL, cTemp, MAX_PATH);  
	spzRealDirEnd = strrchr(cTemp, '\\');
	*spzRealDirEnd = '\0';

	strModulePath = (CString)cTemp;

	return strModulePath;	
}

HRESULT Bluestacks::CreateShortCuttor(LPCSTR szSrcFile, LPSTR szLnkFile, LPSTR szArgument, LPSTR szDesc, LPCSTR szIcoFile) 
{ 
	HRESULT hr; 
	IShellLink *pSl; 
	IPersistFile *pPf; 
	WCHAR wszLnkFile[MAX_PATH]={0,}; 
	 
	hr=CoCreateInstance(CLSID_ShellLink, NULL, CLSCTX_INPROC_SERVER, IID_IShellLink, (void **)&pSl); 
	if (FAILED(hr)) 
		return E_FAIL; 

	pSl->SetPath(szSrcFile); 
	pSl->SetDescription(szDesc); 
	pSl->SetArguments(szArgument); 
	
	pSl->SetIconLocation(szIcoFile, 0);
 
	hr=pSl->QueryInterface(IID_IPersistFile, (void **) &pPf); 
	if (FAILED(hr)) 
	{
		pSl->Release(); 
		return E_FAIL; 
	} 

	MultiByteToWideChar(CP_ACP, 0, szLnkFile, -1, wszLnkFile, MAX_PATH); 
	hr=pPf->Save(wszLnkFile, TRUE); 

	pPf->Release(); 
	pSl->Release(); 
	return hr; 
} 

CString Bluestacks::IsExistDataDir(LPCTSTR szKey)
{	
	char szBuffer [MAX_PATH + 1] = {0, };
	DWORD dwCount = 0;

	CRegKey regKey;

	if (regKey.Open(HKEY_LOCAL_MACHINE, _T("SOFTWARE\\BlueStacks"), KEY_READ) != ERROR_SUCCESS)
	{
		regKey.Close();
		return FALSE;
	}

	dwCount = REG_BUFF_SIZE + 1;
	
	if (ERROR_SUCCESS != regKey.QueryStringValue(szKey, szBuffer, &dwCount))
	{
		regKey.Close();
		return FALSE;
	}

	regKey.Close();

	return (CString)szBuffer;
}

CString Bluestacks::GetJsonFilePath()
{	
	CString strFileFullPath;
	strFileFullPath.Format("%s\\UserData\\Gadget\\apps.json", IsExistDataDir("DataDir"));
	if(!::PathFileExists(strFileFullPath))
		return ""; 

	wstring content;
	CFile file;
	file.Open(strFileFullPath, CFile::modeRead);
	ReadFile::readFile(file, content);
	file.Close();

	CString strTotalContent(content.c_str());
	return strTotalContent;	
}

BOOL Bluestacks::GetSplitAppsTotal(CString strInput, CStringArray& strResultArr)
{
	BOOL bFind = FALSE;

	strResultArr.RemoveAll();
	CString strOutput(_T(""));
	int nCountLeft = strInput.Find(_T("{"));
	int nCountRight = strInput.Find(_T("}"));
	while(nCountRight>nCountLeft)
	{
		strOutput = strInput.Mid(nCountLeft+1, nCountRight - nCountLeft - 1);
		strResultArr.Add(strOutput);
		strInput = strInput.Mid(nCountRight+1);
		nCountLeft = strInput.Find(_T("{"));
		nCountRight = strInput.Find(_T("}"));
		bFind = TRUE;
	}	
	return bFind;
}

BOOL Bluestacks::GetSplitAppsPart(CString strInput, CStringArray& strResultArr)
{
	BOOL bFind = FALSE;

	strResultArr.RemoveAll();
	int nCount = strInput.Find(_T(","));
	int nLength = 0;	
	CString strPrefix, strNext;
	strNext = strInput;
	CString strTmp;
	while(nCount!=-1)
	{
		strPrefix = strNext.Left(nCount);
		strResultArr.Add(strPrefix);
		nLength = strNext.GetLength();
		strTmp = strNext.Right(nLength - nCount - 1);
		strNext = strTmp;
		nCount = strNext.Find(_T(","));
		bFind = TRUE;
	}
	if(strNext != "")strResultArr.Add(strNext);
	return bFind;
}

void Bluestacks::GetSplitString(CString strInput, CString &strName, CString &strValue)
{
	int nCount = strInput.Find(_T(":"));
	int nLength = strInput.GetLength();
	strName = strInput.Left(nCount);
	strValue = strInput.Right(nLength - nCount - 1);

	strName.Replace("\"","");
	strName.TrimLeft();
	strName.TrimRight();

	strValue.Replace("\"","");
	strValue.TrimLeft();
	strValue.TrimRight();
}

BOOL Bluestacks::GetSplitParams(CString strInput, CStringArray& strResultArr)
{
	BOOL bFind = FALSE;

	strResultArr.RemoveAll();
	int nCount = strInput.Find(_T(";"));
	int nLength = 0;	
	CString strPrefix, strNext;
	strNext = strInput;
	CString strTmp;
	while(nCount!=-1)
	{
		strPrefix = strNext.Left(nCount);
		strResultArr.Add(strPrefix);
		nLength = strNext.GetLength();
		strTmp = strNext.Right(nLength - nCount - 1);
		strNext = strTmp;
		nCount = strNext.Find(_T(";"));
	}
	if(strNext != "")strResultArr.Add(strNext);

	if(strResultArr.GetCount() == 3)
		bFind = TRUE;

	return bFind;
}

void Bluestacks::Addat(CString strname,CString strimg,CString strpackage,CString stractivity,CString strsystem,CString strversion, CString strappstore)
{
	JsonApp AppInfos;
	AppInfos.name = strname;
	AppInfos.img = strimg;
	AppInfos.package = strpackage;
	AppInfos.activity = stractivity;
	AppInfos.system = strsystem;
	AppInfos.version = strversion;
	AppInfos.appstore = strappstore;
	m_vInfo.push_back(AppInfos);
}

BOOL Bluestacks::GetSearchApp(CString strpackage, CString stractivity)
{	
	BOOL bFind = FALSE;
	for (int i = 0; i < m_vInfo.size(); i++)
	{		
		if( lstrcmpi(strpackage, m_vInfo[i].package) && lstrcmpi(stractivity, m_vInfo[i].activity))
		{
			bFind = TRUE;
			break;
		}
	}
	return bFind;
}

BOOL Bluestacks::GetSearchApp(CString strname)
{	
	BOOL bFind = FALSE;
	for (int i = 0; i < m_vInfo.size(); i++)
	{
		if( lstrcmpi(strname, m_vInfo[i].name) )
		{
			bFind = TRUE;
			break;
		}
	}
	return bFind;
}

BOOL Bluestacks::GetSearchApp(CString strname, CString &strpackage, CString &stractivity)
{	
	BOOL bFind = FALSE;
	CString strTemp;
	int nFind = 0;
	for (int i = 0; i < m_vInfo.size(); i++)
	{
		strname.MakeLower();
		strTemp = m_vInfo[i].name;
		strTemp.MakeLower();

		if(strTemp.Find(strname) != -1)
		{
			strpackage = m_vInfo[i].package;
			stractivity = m_vInfo[i].activity;
			bFind = TRUE;
			break;
		}
	}
	return bFind;
}

void Bluestacks::DeleteAllMemory()
{
	m_vInfo.clear();
}

BOOL Bluestacks::ReadJson()
{
	CString strInput = GetJsonFilePath();
	if(strInput == "")
		return FALSE;

	CString strName;
	CString strValue;
	CStringArray strTotalApp;
	CStringArray strDetailApp;

	if(GetSplitAppsTotal(strInput, strTotalApp))
	{		
		for(int i=0; i < strTotalApp.GetCount(); i++)
		{			
			if( GetSplitAppsPart(strTotalApp.GetAt(i), strDetailApp) )
			{
				JsonApp AppInfos;

				for(int k=0; k < strDetailApp.GetCount(); k++)
				{				
					GetSplitString(strDetailApp.GetAt(k), strName, strValue);

					if(lstrcmpi(strName, "name") == 0)
						AppInfos.name = strValue;
					else if(lstrcmpi(strName, "img") == 0)
						AppInfos.img = strValue;
					else if(lstrcmpi(strName, "package") == 0)
						AppInfos.package = strValue;
					else if(lstrcmpi(strName, "activity") == 0)
						AppInfos.activity = strValue;
					else if(lstrcmpi(strName, "system") == 0)
						AppInfos.system = strValue;
					else if(lstrcmpi(strName, "version") == 0)
						AppInfos.version = strValue;
					else if(lstrcmpi(strName, "appstore") == 0)
						AppInfos.appstore = strValue;
				}	

				m_vInfo.push_back(AppInfos);
			}
		}
	}	

	return TRUE;
}

BOOL Bluestacks::CreateDirectories(LPCTSTR lpszDirectory)
{
	TCHAR szDir[MAX_PATH+1];
	lstrcpy(szDir, lpszDirectory);

	TCHAR szPath[MAX_PATH+1];
	PathCanonicalize( szPath, szDir );
		
	PathAddBackslash(szPath);

	LPTSTR	lpszSubPath = PathSkipRoot(szPath);
	if ( lpszSubPath == NULL )
		lpszSubPath = szPath;

	LPTSTR	lpszFind;
	TCHAR	szPrefix[MAX_PATH+1];
	
	while ( (lpszFind = _tcschr(lpszSubPath, TCHAR('\\'))) != NULL )
	{
		lstrcpyn(szPrefix, szPath, lpszFind - szPath + 1);
		if ( !CreateDirectory( szPrefix, NULL) )
		{
			if ( ERROR_ALREADY_EXISTS != GetLastError() )
				return FALSE;
		}

		lpszSubPath = _tcsinc(lpszFind);
	}
	return TRUE;
}

void Bluestacks::CreateAppIconList()
{	
	CString strIconName = "";
	CString strIconPath = "";
	CString DataDir = IsExistDataDir("DataDir");
	
    CString strCreateIcon;
	strCreateIcon.Format("%s\\UserData\\Library\\Icons", DataDir);
	int nExist = _access(strCreateIcon, 0);
	if (nExist == -1)
		CreateDirectories(strCreateIcon);

	for (int i = 0; i < m_vInfo.size(); i++)
	{
		char szExePath[MAX_PATH], szLnkPath[MAX_PATH], szArg[MAX_PATH], szExeName[MAX_PATH];
		
		if(m_vInfo[i].img != "")
		{
			int pos = m_vInfo[i].img.ReverseFind('.');
			if(pos >= 0)
				strIconName = m_vInfo[i].img.Left(pos);
			
			strIconPath.Format("%s\\UserData\\Library\\Icons\\%s.ico", DataDir, strIconName);
			if(::PathFileExists(strIconPath))
			{				
				sprintf(szExePath, "%s\\DaeseongLib.exe", GetModulePath());
				sprintf(szLnkPath, "%s\\%s.lnk", GetModulePath(), m_vInfo[i].name);
				sprintf(szArg, "%s;%s;%s", m_vInfo[i].name, m_vInfo[i].package, m_vInfo[i].activity);
				sprintf(szExeName, "%s", m_vInfo[i].name);
				CreateShortCuttor(szExePath, szLnkPath, szArg, szExeName, strIconPath); 		
			}
		}
	}	
}

void Bluestacks::CreateAppIconList(CString strExePath, CString strLnkPath, CString strArg, CString strExeName, CString strIconPath)
{
	char szExePath[MAX_PATH], szLnkPath[MAX_PATH], szArg[MAX_PATH], szExeName[MAX_PATH], szIconPath[MAX_PATH];
	sprintf(szExePath, "%s", strExePath);
	sprintf(szLnkPath, "%s", strLnkPath);
	sprintf(szArg, "%s", strArg);
	sprintf(szExeName, "%s", strExeName);
	sprintf(szIconPath, "%s", strIconPath);
	CreateShortCuttor(szExePath, szLnkPath, szArg, szExeName, szIconPath); 	
}