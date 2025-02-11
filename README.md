# Discord Message Notification GitHub Action

## Overview
This GitHub Action allows you to send a message to a Discord channel using a webhook. It takes in the webhook ID, webhook token, and message content as inputs and sends the message to the specified Discord channel.

## Inputs
| Name           | Description                         | Required |
|-------------- |---------------------------------|----------|
| `webhook_id`  | The Discord webhook ID         | Yes      |
| `webhook_token` | The Discord webhook token     | Yes      |
| `content`     | The message content (max 2000 characters) | Yes      |

## Outputs
| Name     | Description                       |
|----------|---------------------------------|
| `status` | Status of the Discord message   |

