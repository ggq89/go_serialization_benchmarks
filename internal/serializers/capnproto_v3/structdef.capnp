@0x852d6ba8634d7bb6;

using Go = import "/go.capnp";
$Go.package("capnprotov3");
$Go.import(".");

struct CapnpV3 {
  name     @0   :Text;
  birthDay @1   :Int64;
  phone    @2   :Text;
  siblings @3   :Int32;
  spouse   @4   :Bool;
  money    @5   :Float64;
}
