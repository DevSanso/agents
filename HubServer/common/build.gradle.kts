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
    implementation("org.apache.logging.log4j:log4j-core:2.23.0")
    implementation("com.fasterxml.jackson.dataformat:jackson-dataformat-toml:2.16.0")
    runtimeOnly("com.fasterxml.jackson.core:jackson-databind:2.16.1")
    implementation(kotlin("reflect"))
}

tasks.test {
    useJUnitPlatform()
    sourceSets {
        test {
            output.setResourcesDir("build/resources/test")
        }
    }
}