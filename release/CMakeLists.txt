cmake_minimum_required(VERSION 3.12)
project(ia2 )

if(CMAKE_SOURCE_DIR STREQUAL CMAKE_BINARY_DIR OR EXISTS "${CMAKE_BINARY_DIR}/CMakeLists.txt")
  message(FATAL_ERROR "In-source builds are not allowed.")
endif()

find_program(FOUND_GOLANG go)

if(NOT FOUND_GOLANG)
  message(FATAL_ERROR "go is required")
endif()

execute_process(
  COMMAND go run tools/concat.go "${CMAKE_SOURCE_DIR}/lib/IAccessible2/api" "${CMAKE_SOURCE_DIR}/src/ia2.idl"
  WORKING_DIRECTORY "${CMAKE_CURRENT_LIST_DIR}"
)

if(NOT WIN32)
  message(FATAL_ERROR "This project targets Windows.")
endif()

execute_process(
  COMMAND midl ia2.idl
  WORKING_DIRECTORY "${CMAKE_CURRENT_LIST_DIR}/src"
)

file(REMOVE "${CMAKE_SOURCE_DIR}/src/dlldata.c")
file(GLOB HEADERS "${CMAKE_SOURCE_DIR}/src/*.h")

foreach(HEADER ${HEADERS})
  file(COPY "${HEADER}" DESTINATION "${CMAKE_SOURCE_DIR}/include/ia2")
endforeach()

add_subdirectory(src)
