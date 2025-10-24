<p align="center">
    <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png" align="center" width="30%">
</p>
<p align="center"><h1 align="center">GOLANG-TUTORIAL</h1></p>
<p align="center">
 <em><code>❯ A comprehensive beginner-to-intermediate Go programming tutorial collection</code></em>
</p>
<p align="center">
 <img src="https://img.shields.io/github/license/MayurakshaSikdar/golang-tutorial?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
 <img src="https://img.shields.io/github/last-commit/MayurakshaSikdar/golang-tutorial?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
 <img src="https://img.shields.io/github/languages/top/MayurakshaSikdar/golang-tutorial?style=default&color=0080ff" alt="repo-top-language">
 <img src="https://img.shields.io/github/languages/count/MayurakshaSikdar/golang-tutorial?style=default&color=0080ff" alt="repo-language-count">
</p>
<p align="center"><!-- default option, no dependency badges. -->
</p>
<p align="center">
 <!-- default option, no dependency badges. -->
</p>
<br>

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Project Structure](#project-structure)
  - [Project Index](#project-index)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
- [Acknowledgments](#acknowledgments)

---

## Overview

<p>A comprehensive beginner-to-intermediate guide to Go programming. Features code examples, concurrency patterns, REST API development, testing practices, and mini-projects. Ideal for developers starting with Go or brushing up on modern backend development skills.</p>

---

## Features

<li>Progressive learning path from basics to intermediate concepts</li>
<li>Hands-on projects including REST API and text-to-speech application</li>
<li>Practical examples of Go concurrency patterns and best practices</li>
<li>Well-structured code following Go project layout standards</li>
<li>Ready-to-run examples with proper error handling and testing</li>

---

## Project Structure

```sh
└── golang-tutorial/
    ├── README.md
    ├── basics
    │   └── basics.go
    ├── intermmediate
    │   ├── intermmediate1.go
    │   ├── intermmediate2.go
    │   ├── intermmediate3.go
    │   ├── intermmediate4.go
    │   └── intermmediate5.go
    └── projects
        ├── project-api
        └── project-text-speech
```

### Project Index

<details open>
   <summary><b><code>GOLANG-TUTORIAL/</code></b></summary>
   <details style="margin-left: 20px">
      <!-- projects Submodule --> 
      <summary><b>projects</b></summary>
      <blockquote>
         <details>
            <summary style="margin-left: 20px"><b>project-api</b></summary>
            <blockquote>
               <table>
                  <tr>
                     <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-api/go.mod'>go.mod</a></b></td>
                     <td><code>❯ Go module definition for the API project</code></td>
                  </tr>
                  <tr>
                     <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-api/go.sum'>go.sum</a></b></td>
                     <td><code>❯ Dependency checksums for reproducible builds</code></td>
                  </tr>
               </table>
               <details>
                  <summary><b>cmd</b></summary>
                  <blockquote>
                     <details>
                        <summary><b>project-api</b></summary>
                        <blockquote>
                           <table>
                              <tr>
                                 <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-api/cmd/project-api/main.go'>main.go</a></b></td>
                                 <td><code>❯ Application entry point with server initialization</code></td>
                              </tr>
                           </table>
                        </blockquote>
                     </details>
                  </blockquote>
               </details>
               <details>
                  <summary><b>internal</b></summary>
                  <blockquote>
                     <details>
                        <summary><b>types</b></summary>
                        <blockquote>
                           <table>
                              <tr>
                                 <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-api/internal/types/types.go'>types.go</a></b></td>
                                 <td><code>❯ Data structures and domain type definitions</code></td>
                              </tr>
                           </table>
                        </blockquote>
                     </details>
                     <details>
                        <summary><b>http</b></summary>
                        <blockquote>
                           <details>
                              <summary><b>handlers</b></summary>
                              <blockquote>
                                 <details>
                                    <summary><b>student</b></summary>
                                    <blockquote>
                                       <table>
                                          <tr>
                                             <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-api/internal/http/handlers/student/student.go'>student.go</a></b></td>
                                             <td><code>❯ HTTP handlers for student resource operations</code></td>
                                          </tr>
                                       </table>
                                    </blockquote>
                                 </details>
                              </blockquote>
                           </details>
                        </blockquote>
                     </details>
                     <details>
                        <summary><b>config</b></summary>
                        <blockquote>
                           <table>
                              <tr>
                                 <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-api/internal/config/config.go'>config.go</a></b></td>
                                 <td><code>❯ Configuration management and environment setup</code></td>
                              </tr>
                           </table>
                        </blockquote>
                     </details>
                     <details>
                        <summary><b>utils</b></summary>
                        <blockquote>
                           <details>
                              <summary><b>response</b></summary>
                              <blockquote>
                                 <table>
                                    <tr>
                                       <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-api/internal/utils/response/response.go'>response.go</a></b></td>
                                       <td><code>�️❯ Standardized HTTP response utilities</code></td>
                                    </tr>
                                 </table>
                              </blockquote>
                           </details>
                        </blockquote>
                     </details>
                  </blockquote>
               </details>
            </blockquote>
         </details>
         <details style="margin-left: 20px">
            <summary><b>project-text-speech</b></summary>
            <blockquote>
               <table>
                  <tr>
                     <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-text-speech/main.go'>main.go</a></b></td>
                     <td><code>❯ Main entry point for text-to-speech application</code></td>
                  </tr>
                  <tr>
                     <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-text-speech/go.mod'>go.mod</a></b></td>
                     <td><code>❯ Go module definition for text-to-speech project</code></td>
                  </tr>
                  <tr>
                     <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-text-speech/go.sum'>go.sum</a></b></td>
                     <td><code>❯ Dependency checksums for text-to-speech project</code></td>
                  </tr>
               </table>
               <details>
                  <summary><b>speech</b></summary>
                  <blockquote>
                     <table>
                        <tr>
                           <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-text-speech/speech/speech.go'>speech.go</a></b></td>
                           <td><code>❯ Core text-to-speech conversion logic</code></td>
                        </tr>
                     </table>
                  </blockquote>
               </details>
               <details>
                  <summary><b>cmd</b></summary>
                  <blockquote>
                     <table>
                        <tr>
                           <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/projects/project-text-speech/cmd/cmd.go'>cmd.go</a></b></td>
                           <td><code>❯ Command-line interface and argument parsing</code></td>
                        </tr>
                     </table>
                  </blockquote>
               </details>
            </blockquote>
         </details>
      </blockquote>
   </details>
   <details style="margin-left: 20px">
      <!-- basics Submodule --> 
      <summary><b>basics</b></summary>
      <blockquote>
         <table>
            <tr>
               <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/basics/basics.go'>basics.go</a></b></td>
               <td><code>❯ Fundamental Go concepts: variables, functions, control structures</code></td>
            </tr>
         </table>
      </blockquote>
   </details>
   <details style="margin-left: 20px">
      <!-- intermmediate Submodule --> 
      <summary><b>intermmediate</b></summary>
      <blockquote>
         <table>
            <tr>
               <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/intermmediate/intermmediate1.go'>intermmediate1.go</a></b></td>
               <td><code>❯ Advanced topics: interfaces and type assertions</code></td>
            </tr>
            <tr>
               <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/intermmediate/intermmediate2.go'>intermmediate2.go</a></b></td>
               <td><code>❯ Concurrency patterns with goroutines and channels</code></td>
            </tr>
            <tr>
               <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/intermmediate/intermmediate5.go'>intermmediate5.go</a></b></td>
               <td><code>❯ Error handling strategies and custom error types</code></td>
            </tr>
            <tr>
               <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/intermmediate/intermmediate4.go'>intermmediate4.go</a></b></td>
               <td><code>❯ Working with JSON and struct tags</code></td>
            </tr>
            <tr>
               <td><b><a href='https://github.com/MayurakshaSikdar/golang-tutorial/blob/master/intermmediate/intermmediate3.go'>intermmediate3.go</a></b></td>
               <td><code>❯ File I/O operations and buffer management</code></td>
            </tr>
         </table>
      </blockquote>
   </details>
</details>

---

## Getting Started

### Prerequisites

Before getting started with golang-tutorial, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go 1.16 or higher
- **Package Manager:** Go modules

### Installation

Install golang-tutorial using one of the following methods:

**Build from source:**

1. Clone the golang-tutorial repository:

```sh
❯ git clone https://github.com/MayurakshaSikdar/golang-tutorial
```

2. Navigate to the project directory:

```sh
❯ cd golang-tutorial
```

3. Install the project dependencies:

**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go build
```

### Usage

Run golang-tutorial using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run {entrypoint}
```

## Acknowledgments

- Thanks to the Go community for excellent documentation and learning resources from [CoderGyan](https://codersgyan.com) ft. [Rakesh](https://www.linkedin.com/in/codersgyan/)
- Inspired by real-world backend development patterns and best practices
- Based on proven Go project structure standards and conventions

---
