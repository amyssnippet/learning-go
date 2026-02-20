# Modules and Dependencies


- ## `go mod init`:

it is used to initialize a new go module by creating a `go.mod` file with specified module path like repo links. It marks directory as module root and enables module based dependency management. it is the first step for any new go project.

- ## `go mod tidy`:

it checks the entire codebase, updates new dependency and removes unused ones and also updates `go.sum` file with checksums. it is essential for maintaining clean dependency management and ensuring reproducible builds before prod deploys.

- ## `go mod vendor`:

creates a vendor directory with dependency copies for bundling with source code. it ensures build works without internet access. useful for deployment, air gapped environments, and complete control over dependency availability



# Packages

Packages are fundamental unit of code organisation in go. group related types, functions and variables defined at the top of the file with `package` keyword. it enable modularity, reusability and namespace management.


# Package Import RUles

key rules include:

- no circular imports
- main package for executables
- lowercase package names
- exported identifiers start with capitals
- import paths are unique identifiers



# Third Party Packages

in go, we can import external libraries by `go get package-url`, which updates the `go.mod`. 



# Publishing Modules

Share Go code through version control systems using semantic versioning tags. Go proxy system automatically discovers and serves modules. Follow Go conventions, maintain documentation, and ensure backward compatibility to contribute to the ecosystem.