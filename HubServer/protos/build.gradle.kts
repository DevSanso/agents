plugins {
    kotlin("jvm")
    id("com.google.protobuf").version("0.9.4")
}

group = "devsanso.github.io.HubServer"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    testImplementation("org.jetbrains.kotlin:kotlin-test")
    implementation("com.google.protobuf:protobuf-kotlin:3.25.3")
    runtimeOnly("com.google.protobuf:protobuf-gradle-plugin:0.9.4")

}

sourceSets {
    main {
        proto {
            srcDir("../../protobuf")
        }
    }
}

tasks.test {
    useJUnitPlatform()
}

