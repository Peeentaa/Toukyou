#include <iostream>
#include <sqlite3.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
    std::string databasePath = argv[2];

    sqlite3 *db;
    int rc = sqlite3_open(databasePath.c_str(), &db);
    if (rc) {
        fprintf(stderr, "Can't open database: %s\n", sqlite3_errmsg(db));
        return 1;
    } else {
        fprintf(stderr, "Opened database successfully\n");
    }

    sqlite3_close(db);
    return 0;
}
