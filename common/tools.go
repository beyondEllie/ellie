package common

import "github.com/tacheraSasi/ellie/types"



var Tools []types.DevTool = []types.DevTool{
	{
		Name:          "Git",
		Description:   "Version control system",
		CheckCmd:      "git --version",
		DefaultInstall: true,
		Install: map[string]string{
			"mac":    "brew install git",
			"linux":  "sudo apt-get install git -y",
			"windows": "choco install git -y",
		},
		Configure: map[string]string{
			"common": "git config --global core.autocrlf input && git config --global init.defaultBranch main",
		},
	},
	{
		Name:          "Node.js",
		Description:   "JavaScript runtime",
		CheckCmd:      "node --version",
		DefaultInstall: true,
		Install: map[string]string{
			"mac":    "brew install node",
			"linux":  "curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash - && sudo apt-get install -y nodejs",
			"windows": "choco install nodejs-lts",
		},
	},
	{
		Name:          "Docker",
		Description:   "Containerization platform",
		CheckCmd:      "docker --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install --cask docker",
			"linux":  "curl -fsSL https://get.docker.com | sh",
			"windows": "choco install docker-desktop",
		},
	},
	{
		Name:          "Go",
		Description:   "Go programming language",
		CheckCmd:      "go version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install go",
			"linux":  "sudo apt install golang -y",
			"windows": "choco install golang",
		},
	},
	{
		Name:          "Python",
		Description:   "Python programming language",
		CheckCmd:      "python3 --version",
		DefaultInstall: true,
		Install: map[string]string{
			"mac":    "brew install python",
			"linux":  "sudo apt install python3 -y",
			"windows": "choco install python",
		},
	},
	{
		Name:          "VS Code",
		Description:   "Code editor",
		CheckCmd:      "code --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install --cask visual-studio-code",
			"linux":  "sudo snap install --classic code",
			"windows": "choco install vscode",
		},
	},
	{
		Name:          "Yarn",
		Description:   "Modern package manager",
		CheckCmd:      "yarn --version",
		DefaultInstall: true,
		Install: map[string]string{
			"mac":    "brew install yarn",
			"linux":  "npm install -g yarn",
			"windows": "choco install yarn",
		},
	},
	{
		Name:          "PostgreSQL",
		Description:   "Relational database",
		CheckCmd:      "psql --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install postgresql",
			"linux":  "sudo apt install postgresql postgresql-contrib -y",
			"windows": "choco install postgresql",
		},
		Configure: map[string]string{
			"linux": "sudo systemctl enable postgresql && sudo systemctl start postgresql",
			"mac":   "brew services start postgresql",
		},
	},
	{
		Name:          "Redis",
		Description:   "In-memory data store",
		CheckCmd:      "redis-cli --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install redis",
			"linux":  "sudo apt install redis -y",
			"windows": "choco install redis",
		},
		Configure: map[string]string{
			"linux": "sudo systemctl enable redis && sudo systemctl start redis",
			"mac":   "brew services start redis",
		},
	},
	{
		Name:          "AWS CLI",
		Description:   "Amazon Web Services CLI",
		CheckCmd:      "aws --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install awscli",
			"linux":  "curl 'https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip' -o awscliv2.zip && unzip awscliv2.zip && sudo ./aws/install",
			"windows": "choco install awscli",
		},
	},
	{
		Name:          "Terraform",
		Description:   "Infrastructure as code tool",
		CheckCmd:      "terraform --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install terraform",
			"linux":  "sudo apt install terraform -y",
			"windows": "choco install terraform",
		},
	},
	{
		Name:          "kubectl",
		Description:   "Kubernetes cluster manager",
		CheckCmd:      "kubectl version --client",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install kubectl",
			"linux":  "sudo apt install kubectl -y",
			"windows": "choco install kubernetes-cli",
		},
	},
	{
		Name:          "Helm",
		Description:   "Kubernetes package manager",
		CheckCmd:      "helm version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install helm",
			"linux":  "curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash",
			"windows": "choco install kubernetes-helm",
		},
	},
	{
		Name:          "NGINX",
		Description:   "Web server and reverse proxy",
		CheckCmd:      "nginx -v",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install nginx",
			"linux":  "sudo apt install nginx -y",
			"windows": "choco install nginx",
		},
	},
	{
		Name:          "GitHub CLI",
		Description:   "GitHub command-line tool",
		CheckCmd:      "gh --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install gh",
			"linux":  "sudo apt install gh -y",
			"windows": "choco install gh",
		},
	},
	{
		Name:          ".NET SDK",
		Description:   ".NET development platform",
		CheckCmd:      "dotnet --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install dotnet",
			"linux":  "wget https://dot.net/v1/dotnet-install.sh -O dotnet-install.sh && chmod +x ./dotnet-install.sh && ./dotnet-install.sh",
			"windows": "choco install dotnet-sdk",
		},
	},
	{
		Name:          "Java",
		Description:   "OpenJDK development kit",
		CheckCmd:      "javac --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install openjdk",
			"linux":  "sudo apt install openjdk-17-jdk -y",
			"windows": "choco install openjdk",
		},
	},
	{
		Name:          "PHP",
		Description:   "PHP runtime",
		CheckCmd:      "php --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install php",
			"linux":  "sudo apt install php -y",
			"windows": "choco install php",
		},
	},
	{
		Name:          "Ansible",
		Description:   "Configuration management",
		CheckCmd:      "ansible --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install ansible",
			"linux":  "sudo apt install ansible -y",
			"windows": "choco install ansible",
		},
	},
	{
		Name:          "Vagrant",
		Description:   "Virtual machine manager",
		CheckCmd:      "vagrant --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":    "brew install vagrant",
			"linux":  "sudo apt install vagrant -y",
			"windows": "choco install vagrant",
		},
	},
	{
		Name:          "Prettier",
		Description:   "Code formatter",
		CheckCmd:      "prettier --version",
		DefaultInstall: true,
		Install: map[string]string{
			"common": "npm install -g prettier",
		},
	},
	{
		Name:          "ESLint",
		Description:   "JavaScript linter",
		CheckCmd:      "eslint --version",
		DefaultInstall: true,
		Install: map[string]string{
			"common": "npm install -g eslint",
		},
	},
}