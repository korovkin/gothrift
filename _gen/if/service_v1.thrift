include "base_service_v1.thrift"

struct Location {
1: double timestamp_unix_sec;
2: double latitude_degrees;
3: double longitude_degrees;
}

service Blackbox extends base_service_v1.BaseService {
   void ping();
   string get_version();
   void log_location(1: Location loc);
}
