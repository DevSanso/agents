plugins {
    kotlin("jvm") version "1.9.21"
    java
}

group = "devsanso.github.io.HubServer"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    testImplementation("org.jetbrains.kotlin:kotlin-test")
}

tasks.test {
    useJUnitPlatform()
}
