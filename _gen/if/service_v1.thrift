include "base_service_v1.thrift"

service Blackbox extends base_service_v1.BaseService 
{
   void ping();
   string get_version();
}
