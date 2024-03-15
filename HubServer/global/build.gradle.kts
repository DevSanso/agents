plugins {
    kotlin("jvm")
}

group = "devsanso.github.io.HubServer"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    testImplementation("org.jetbrains.kotlin:kotlin-test")
    compileOnly(project(":common"))
    compileOnly(project(":protos"))
    compileOnly("com.google.protobuf:protobuf-kotlin:3.25.3")
}

tasks.test {
    useJUnitPlatform()

    dependencies {
        implementation(project(":protos"))
        implementation("com.google.protobuf:protobuf-kotlin:3.25.3")
    }
}