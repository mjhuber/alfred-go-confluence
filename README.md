# alfred-go-confluence

an alfred workflow to search confluence

## Setup

1. Go to the releases page and download/import the workflow into alred.
2. Edit the workflow and make sure the following variables are set:

* `CONFLUENCE_BASEURL` (ex: `https://<your-company>.atlassian.net`)
* `CONFLUENCE_TOKEN`: Your confluence access token.  You can generate a token in your [Atlassian Account Settings](https://id.atlassian.com/manage-profile/security/api-tokens)
* `CONFLUENCE_USERNAME`: Your conflunece username

## FAQ

**Q: I get the error `“alfred-go-confluence” cannot be opened because the developer cannot be verified.`.  How can i fix it?

**A: Thats an error gatekeeper returns when a binary isn't signed by an apple certificate.  To fix that follow these steps:

  1. Press cancel on the promp you received.
  2. Go to preferences and select `Security & Privacy`.  On the general tab, make sure "Allow apps download from: App Store and identified developers" is selected.  Beneath that you should see the text `alfred-go-confluence was blocked from use becuase it is not from an identified developer`.  Press the `Allow Anyway` button and it try using the workflow again.
