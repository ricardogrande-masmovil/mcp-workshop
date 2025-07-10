# How to configure MCPs in VSCode

The best way to understand all configuration possibilities is to read [the docs](https://code.visualstudio.com/docs/copilot/chat/mcp-servers)! 

# Some interesting MCPs

## Github official MCP
The official MCP is available at [https://github.com/github/github-mcp-server](https://github.com/github/github-mcp-server)

### Using OAuth
```
{
  "servers": {
    "github": {
      "type": "http",
      "url": "https://api.githubcopilot.com/mcp/"
    }
  }
}
```

### Using a personal access token
```
{
  "servers": {
    "github": {
      "type": "http",
      "url": "https://api.githubcopilot.com/mcp/",
      "headers": {
        "Authorization": "Bearer ${input:github_mcp_pat}"
      }
    }
  },
  "inputs": [
    {
      "type": "promptString",
      "id": "github_mcp_pat",
      "description": "GitHub Personal Access Token",
      "password": true
    }
  ]
}
```

## Atlassian Cloud Official MCP
The official MCP is available at [https://community.atlassian.com/forums/Atlassian-Platform-articles/Using-the-Atlassian-Remote-MCP-Server-beta/ba-p/3005104](https://community.atlassian.com/forums/Atlassian-Platform-articles/Using-the-Atlassian-Remote-MCP-Server-beta/ba-p/3005104)

### Using OAuth
```
{
  "servers": {
    "atlassian": {
      "type": "http",
      "url": "https://mcp.atlassian.com/v1/sse"
    }
  }
}
```

## Atlassian (Server/Legacy) MCP
Caution! This is not an official MCP server, but while we migrate Jira to Atlassian Cloud, there is no other option. 
We will use PAT to authenticate, so be aware that the scopes of the token are the same as the permissions of the user.
Any action will be signed by the user owning the token.

Refer to the github repository for more information: [https://github.com/sooperset/mcp-atlassian](https://github.com/sooperset/mcp-atlassian)

### Using a personal access token
```
{
    "servers": {
        "mcp-atlassian": {
            "command": "docker",
            "args": [
                "run",
                "--rm",
                "-i",
                "-e", "JIRA_URL",
                "-e", "JIRA_PERSONAL_TOKEN",
                "-e", "JIRA_SSL_VERIFY",
                "ghcr.io/sooperset/mcp-atlassian:latest"
            ],
            "env": {
                "JIRA_URL": "https://jiranext.masorange.es",
                "JIRA_PERSONAL_TOKEN": "${input:jira_mcp_pat}",
                "JIRA_SSL_VERIFY": "false"
            }
        }
    },
    "inputs": [
        {
            "id": "jira_mcp_pat",
            "type": "promptString",
            "description": "Enter your JIRA Personal Access Token for MCP integration",
            "password": true
        }
    ],
}
```

## PostgreSQL MCP
The official MCP for read-only access to a PostgreSQL database is available at [https://github.com/modelcontextprotocol/servers-archived/tree/main/src/postgres](https://github.com/modelcontextprotocol/servers-archived/tree/main/src/postgres)

```
{
    "inputs": [
        {
            "type": "promptString",
            "id": "pg_user",
            "description": "PostgreSQL User"
        },
        {
            "type": "promptString",
            "id": "pg_password",
            "description": "PostgreSQL Password",
            "password": true
        },
        {
            "type": "promptString",
            "id": "pg_host",
            "description": "PostgreSQL Host (e.g. localhost)"
        },
        {
            "type": "promptString",
            "id": "pg_port",
            "description": "PostgreSQL Port (e.g. 5432)"
        },
        {
            "type": "promptString",
            "id": "pg_database",
            "description": "PostgreSQL Database Name"
        }
    ],
    "servers": {
        "postgres": {
            "command": "npx",
            "args": [
                "-y",
                "@modelcontextprotocol/server-postgres",
                "postgresql://${input:pg_user}:${input:pg_password}@${input:pg_host}:${input:pg_port}/${input:pg_database}",
            ]
        }
    }
}
```

## Some Inspiration
There are many official MCPs you can use in you daily work, [here are some examples](https://github.com/modelcontextprotocol/servers?tab=readme-ov-file#%EF%B8%8F-official-integrations).

Think about how easy it is to set up MCPs, regardless of the client you are using.
Think about AI Agents connected to unlimited capabilities. Think about your daily processes, I am sure you can find many steps that can be automated with LLMs using MCP servers.
You can even create your own MCP server to expose your own data and processes, and use it in your daily work.