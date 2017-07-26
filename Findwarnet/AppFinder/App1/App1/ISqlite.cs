using System;
using SQLite;

namespace App1
{
    public interface ISqlite
    {
        SQLiteConnection GetConnection();
    }
}
