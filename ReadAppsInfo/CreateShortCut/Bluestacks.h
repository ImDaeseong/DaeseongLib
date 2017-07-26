#pragma once

struct JsonApp
{
	CString name;
	CString img;
	CString package;
	CString activity;
	CString system;
	CString version;
	CString appstore;
};

class Bluestacks
{
public:
	Bluestacks();
	virtual ~Bluestacks();

public:
	
	BOOL ReadJson();
	CString GetJsonFilePath();
	CString IsExistDataDir(LPCTSTR szKey);
	
	BOOL GetSplitAppsTotal(CString strInput, CStringArray& strResultArr);
	BOOL GetSplitAppsPart(CString strInput, CStringArray& strResultArr);
	void GetSplitString(CString strInput, CString &strName, CString &strValue);
    BOOL GetSplitParams(CString strInput, CStringArray& strResultArr);

	void Addat(CString strname,CString strimg,CString strpackage,CString stractivity,CString strsystem,CString strversion, CString strappstore);
	BOOL GetSearchApp(CString strpackage, CString stractivity);
	BOOL GetSearchApp(CString strname);
	BOOL GetSearchApp(CString strname, CString &strpackage, CString &stractivity);
	void DeleteAllMemory();

	HRESULT CreateShortCuttor(LPCSTR szSrcFile, LPSTR szLnkFile, LPSTR szArgument, LPSTR szDesc, LPCSTR szIcoFile);	
			
	void CreateAppIconList();
	void CreateAppIconList(CString strExePath, CString strLnkPath, CString strArg, CString strExeName, CString strIconPath);

	BOOL CreateDirectories(LPCTSTR lpszDirectory);
	CString GetModulePath();

	std::vector<JsonApp> GetJsonInfo() { return m_vInfo; }

private:
	std::vector<JsonApp> m_vInfo;
};


