package constant

// CommonSuffixes spring组件的变体常量，包含这些名称的spring组件都要写单独的逻辑创建变体
// -core 去除 -core
// -api  去除 -api
// -classic 去除 -classic
// spring-webmvc -- spring_framework
// - 替换成 _
var CommonSuffixes = []string{"-api", "-core", ".core", "-client-core", "-classic", "-api", "-complete", "-full", "-all", "-ex", "-server", ".js", "-handler", "apache-", "-web", "-broker", "-netty", "-plugin", "-web-console", "-main", "-war"}

var OsPkgTypes = []string{"deb", "apk", "rpm", "swid", "alpm", "docker", "oci", "generic", "qpkg", "buildroot", "coreos", "ebuild", "alpine", "alma", "debian", "ubuntu", "amazon", "redhat", "rocky", "arch", "suse", "photon", "microsoft"}
