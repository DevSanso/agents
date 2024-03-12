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
    implementation(project(":global"))
    runtimeOnly("org.postgresql:postgresql:42.7.2")
    implementation("org.mybatis:mybatis:3.5.15")

}

tasks.test {
    useJUnitPlatform()
}