import org.jetbrains.kotlin.cli.jvm.compiler.findMainClass

plugins {
    kotlin("jvm")
    application
}

group = "devsanso.github.io.HubServer"
version = "1.0"
val appMainClassName = "devsanso.github.io.HubServer.MainKt"

repositories {
    mavenCentral()
}

dependencies {
    testImplementation("org.jetbrains.kotlin:kotlin-test")
    implementation("org.jetbrains.kotlin:kotlin-stdlib:1.9.22")
    implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8:1.9.22")
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.8.0")
    compileOnly(project(":common"))
    implementation(project(":global"))
    implementation(project(":sender"))
    implementation(project(":receiver"))
}

application {

    applicationName = "HubServer"
    mainClass.set(appMainClassName)

}

tasks.test {
    useJUnitPlatform()
}

tasks.jar {
    manifest {
        attributes["Main-Class"] = appMainClassName
    }
}

tasks.run {
    dependencies {
        implementation(project(":common"))
    }
}
