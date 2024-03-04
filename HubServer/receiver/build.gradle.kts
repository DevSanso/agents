plugins {
    kotlin("jvm")
}

group = "devsanso.github.io"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    testImplementation("org.jetbrains.kotlin:kotlin-test")
    implementation("redis.clients:jedis:5.1.1")
    compileOnly(project(":common"))
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.8.0")
    runtimeOnly("org.jetbrains.kotlinx:kotlinx-coroutines-core-jvm:1.8.0")
}

tasks.test {
    useJUnitPlatform()
    dependencies {
        implementation(project(":common"))
    }
}