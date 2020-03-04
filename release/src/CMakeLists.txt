file(GLOB HEADERS_IA2"../include/ia2/*.h")
source_group("Header Files\\ia2" FILES ${HEADERS_IA2})

file(GLOB SOURCES "./*.c")
list(APPEND SOURCES ${HEADERS_IA2})

add_library(ia2 ${SOURCES})
target_include_directories(ia2 PUBLIC $<BUILD_INTERFACE:${PROJECT_SOURCE_DIR}/include>)
