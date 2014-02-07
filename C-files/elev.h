




<!DOCTYPE html>
<html>
  <head prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# object: http://ogp.me/ns/object# article: http://ogp.me/ns/article# profile: http://ogp.me/ns/profile#">
    <meta charset='utf-8'>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <title>TTK4145/Project/driver/elev.h at master · klasbo/TTK4145</title>
    <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="GitHub" />
    <link rel="fluid-icon" href="https://github.com/fluidicon.png" title="GitHub" />
    <link rel="apple-touch-icon" sizes="57x57" href="/apple-touch-icon-114.png" />
    <link rel="apple-touch-icon" sizes="114x114" href="/apple-touch-icon-114.png" />
    <link rel="apple-touch-icon" sizes="72x72" href="/apple-touch-icon-144.png" />
    <link rel="apple-touch-icon" sizes="144x144" href="/apple-touch-icon-144.png" />
    <meta property="fb:app_id" content="1401488693436528"/>

      <meta content="@github" name="twitter:site" /><meta content="summary" name="twitter:card" /><meta content="klasbo/TTK4145" name="twitter:title" /><meta content="Contribute to TTK4145 development by creating an account on GitHub." name="twitter:description" /><meta content="https://0.gravatar.com/avatar/15cebd3d1dfcbf51093ec71bb5649e2a?d=https%3A%2F%2Fidenticons.github.com%2Ff4f9cec9fec60d40d51bfdbb5894cc36.png&amp;r=x&amp;s=400" name="twitter:image:src" />
<meta content="GitHub" property="og:site_name" /><meta content="object" property="og:type" /><meta content="https://0.gravatar.com/avatar/15cebd3d1dfcbf51093ec71bb5649e2a?d=https%3A%2F%2Fidenticons.github.com%2Ff4f9cec9fec60d40d51bfdbb5894cc36.png&amp;r=x&amp;s=400" property="og:image" /><meta content="klasbo/TTK4145" property="og:title" /><meta content="https://github.com/klasbo/TTK4145" property="og:url" /><meta content="Contribute to TTK4145 development by creating an account on GitHub." property="og:description" />

    <meta name="hostname" content="github-fe132-cp1-prd.iad.github.net">
    <meta name="ruby" content="ruby 2.1.0p0-github-tcmalloc (87d8860372) [x86_64-linux]">
    <link rel="assets" href="https://github.global.ssl.fastly.net/">
    <link rel="conduit-xhr" href="https://ghconduit.com:25035/">
    <link rel="xhr-socket" href="/_sockets" />
    


    <meta name="msapplication-TileImage" content="/windows-tile.png" />
    <meta name="msapplication-TileColor" content="#ffffff" />
    <meta name="selected-link" value="repo_source" data-pjax-transient />
    <meta content="collector.githubapp.com" name="octolytics-host" /><meta content="collector-cdn.github.com" name="octolytics-script-host" /><meta content="github" name="octolytics-app-id" /><meta content="81F1BB95:53AC:3EA2B3B:52F4D3E3" name="octolytics-dimension-request_id" /><meta content="6430053" name="octolytics-actor-id" /><meta content="hakonBG" name="octolytics-actor-login" /><meta content="143b804ba97286f92de9836672dafefc310d4d773f49d49b3f30d097eb850b00" name="octolytics-actor-hash" />
    

    
    
    <link rel="icon" type="image/x-icon" href="/favicon.ico" />

    <meta content="authenticity_token" name="csrf-param" />
<meta content="SsN7bRjzabuzXQYM/m1vHRQvmUDmCcH7hiygyWP07cM=" name="csrf-token" />

    <link href="https://github.global.ssl.fastly.net/assets/github-f75ad12241fc0d2d7c903870c877da49bf925f5b.css" media="all" rel="stylesheet" type="text/css" />
    <link href="https://github.global.ssl.fastly.net/assets/github2-eedb06b089d2ef792059bc2ebee160f716fa7aab.css" media="all" rel="stylesheet" type="text/css" />
    


      <script src="https://github.global.ssl.fastly.net/assets/frameworks-e8d62aa911c75d1d60662859d52c3cf1232675e6.js" type="text/javascript"></script>
      <script async="async" defer="defer" src="https://github.global.ssl.fastly.net/assets/github-62b15e7c9c7aead7539d1a6b7523410cfd4d33a5.js" type="text/javascript"></script>
      
      <meta http-equiv="x-pjax-version" content="59617e98acce9d59e36d799192b445af">

        <link data-pjax-transient rel='permalink' href='/klasbo/TTK4145/blob/64551ef1eca244838f734d143e8e8b40d256e88d/Project/driver/elev.h'>

  <meta name="description" content="Contribute to TTK4145 development by creating an account on GitHub." />

  <meta content="6306313" name="octolytics-dimension-user_id" /><meta content="klasbo" name="octolytics-dimension-user_login" /><meta content="15593829" name="octolytics-dimension-repository_id" /><meta content="klasbo/TTK4145" name="octolytics-dimension-repository_nwo" /><meta content="true" name="octolytics-dimension-repository_public" /><meta content="false" name="octolytics-dimension-repository_is_fork" /><meta content="15593829" name="octolytics-dimension-repository_network_root_id" /><meta content="klasbo/TTK4145" name="octolytics-dimension-repository_network_root_nwo" />
  <link href="https://github.com/klasbo/TTK4145/commits/master.atom" rel="alternate" title="Recent Commits to TTK4145:master" type="application/atom+xml" />

  </head>


  <body class="logged_in  env-production linux vis-public page-blob">
    <div class="wrapper">
      
      
      
      


      <div class="header header-logged-in true">
  <div class="container clearfix">

    <a class="header-logo-invertocat" href="https://github.com/">
  <span class="mega-octicon octicon-mark-github"></span>
</a>

    
    <a href="/notifications" class="notification-indicator tooltipped downwards" data-gotokey="n" title="You have no unread notifications">
        <span class="mail-status all-read"></span>
</a>

      <div class="command-bar js-command-bar  in-repository">
          <form accept-charset="UTF-8" action="/search" class="command-bar-form" id="top_search_form" method="get">

<input type="text" data-hotkey="/ s" name="q" id="js-command-bar-field" placeholder="Search or type a command" tabindex="1" autocapitalize="off"
    
    data-username="hakonBG"
      data-repo="klasbo/TTK4145"
      data-branch="master"
      data-sha="d85a76d143cf0016166ad1e0e1d0149adc4eae39"
  >

    <input type="hidden" name="nwo" value="klasbo/TTK4145" />

    <div class="select-menu js-menu-container js-select-menu search-context-select-menu">
      <span class="minibutton select-menu-button js-menu-target">
        <span class="js-select-button">This repository</span>
      </span>

      <div class="select-menu-modal-holder js-menu-content js-navigation-container">
        <div class="select-menu-modal">

          <div class="select-menu-item js-navigation-item js-this-repository-navigation-item selected">
            <span class="select-menu-item-icon octicon octicon-check"></span>
            <input type="radio" class="js-search-this-repository" name="search_target" value="repository" checked="checked" />
            <div class="select-menu-item-text js-select-button-text">This repository</div>
          </div> <!-- /.select-menu-item -->

          <div class="select-menu-item js-navigation-item js-all-repositories-navigation-item">
            <span class="select-menu-item-icon octicon octicon-check"></span>
            <input type="radio" name="search_target" value="global" />
            <div class="select-menu-item-text js-select-button-text">All repositories</div>
          </div> <!-- /.select-menu-item -->

        </div>
      </div>
    </div>

  <span class="octicon help tooltipped downwards" title="Show command bar help">
    <span class="octicon octicon-question"></span>
  </span>


  <input type="hidden" name="ref" value="cmdform">

</form>
        <ul class="top-nav">
          <li class="explore"><a href="/explore">Explore</a></li>
            <li><a href="https://gist.github.com">Gist</a></li>
            <li><a href="/blog">Blog</a></li>
          <li><a href="https://help.github.com">Help</a></li>
        </ul>
      </div>

    


  <ul id="user-links">
    <li>
      <a href="/hakonBG" class="name">
        <img alt="hakonBG" height="20" src="https://0.gravatar.com/avatar/dd5b85d4a4740272b321e862738ac77a?d=https%3A%2F%2Fidenticons.github.com%2F47658fe5aed511b671c02813411226c9.png&amp;r=x&amp;s=140" width="20" /> hakonBG
      </a>
    </li>

    <li class="new-menu dropdown-toggle js-menu-container">
      <a href="#" class="js-menu-target tooltipped downwards" title="Create new..." aria-label="Create new...">
        <span class="octicon octicon-plus"></span>
        <span class="dropdown-arrow"></span>
      </a>

      <div class="js-menu-content">
      </div>
    </li>

    <li>
      <a href="/settings/profile" id="account_settings"
        class="tooltipped downwards"
        aria-label="Account settings "
        title="Account settings ">
        <span class="octicon octicon-tools"></span>
      </a>
    </li>
    <li>
      <a class="tooltipped downwards" href="/logout" data-method="post" id="logout" title="Sign out" aria-label="Sign out">
        <span class="octicon octicon-log-out"></span>
      </a>
    </li>

  </ul>



<div class="js-new-dropdown-contents hidden">
  

<ul class="dropdown-menu">
  <li>
    <a href="/new"><span class="octicon octicon-repo-create"></span> New repository</a>
  </li>
  <li>
    <a href="/organizations/new"><span class="octicon octicon-organization"></span> New organization</a>
  </li>


    <li class="section-title">
      <span title="klasbo/TTK4145">This repository</span>
    </li>
      <li>
        <a href="/klasbo/TTK4145/issues/new"><span class="octicon octicon-issue-opened"></span> New issue</a>
      </li>
</ul>

</div>


    
  </div>
</div>

      

      




          <div class="site" itemscope itemtype="http://schema.org/WebPage">
    
    <div class="pagehead repohead instapaper_ignore readability-menu">
      <div class="container">
        

<ul class="pagehead-actions">

    <li class="subscription">
      <form accept-charset="UTF-8" action="/notifications/subscribe" class="js-social-container" data-autosubmit="true" data-remote="true" method="post"><div style="margin:0;padding:0;display:inline"><input name="authenticity_token" type="hidden" value="SsN7bRjzabuzXQYM/m1vHRQvmUDmCcH7hiygyWP07cM=" /></div>  <input id="repository_id" name="repository_id" type="hidden" value="15593829" />

    <div class="select-menu js-menu-container js-select-menu">
      <a class="social-count js-social-count" href="/klasbo/TTK4145/watchers">
        3
      </a>
      <span class="minibutton select-menu-button with-count js-menu-target" role="button" tabindex="0">
        <span class="js-select-button">
          <span class="octicon octicon-eye-watch"></span>
          Watch
        </span>
      </span>

      <div class="select-menu-modal-holder">
        <div class="select-menu-modal subscription-menu-modal js-menu-content">
          <div class="select-menu-header">
            <span class="select-menu-title">Notification status</span>
            <span class="octicon octicon-remove-close js-menu-close"></span>
          </div> <!-- /.select-menu-header -->

          <div class="select-menu-list js-navigation-container" role="menu">

            <div class="select-menu-item js-navigation-item selected" role="menuitem" tabindex="0">
              <span class="select-menu-item-icon octicon octicon-check"></span>
              <div class="select-menu-item-text">
                <input checked="checked" id="do_included" name="do" type="radio" value="included" />
                <h4>Not watching</h4>
                <span class="description">You only receive notifications for conversations in which you participate or are @mentioned.</span>
                <span class="js-select-button-text hidden-select-button-text">
                  <span class="octicon octicon-eye-watch"></span>
                  Watch
                </span>
              </div>
            </div> <!-- /.select-menu-item -->

            <div class="select-menu-item js-navigation-item " role="menuitem" tabindex="0">
              <span class="select-menu-item-icon octicon octicon octicon-check"></span>
              <div class="select-menu-item-text">
                <input id="do_subscribed" name="do" type="radio" value="subscribed" />
                <h4>Watching</h4>
                <span class="description">You receive notifications for all conversations in this repository.</span>
                <span class="js-select-button-text hidden-select-button-text">
                  <span class="octicon octicon-eye-unwatch"></span>
                  Unwatch
                </span>
              </div>
            </div> <!-- /.select-menu-item -->

            <div class="select-menu-item js-navigation-item " role="menuitem" tabindex="0">
              <span class="select-menu-item-icon octicon octicon-check"></span>
              <div class="select-menu-item-text">
                <input id="do_ignore" name="do" type="radio" value="ignore" />
                <h4>Ignoring</h4>
                <span class="description">You do not receive any notifications for conversations in this repository.</span>
                <span class="js-select-button-text hidden-select-button-text">
                  <span class="octicon octicon-mute"></span>
                  Stop ignoring
                </span>
              </div>
            </div> <!-- /.select-menu-item -->

          </div> <!-- /.select-menu-list -->

        </div> <!-- /.select-menu-modal -->
      </div> <!-- /.select-menu-modal-holder -->
    </div> <!-- /.select-menu -->

</form>
    </li>

  <li>
  

  <div class="js-toggler-container js-social-container starring-container ">
    <a href="/klasbo/TTK4145/unstar"
      class="minibutton with-count js-toggler-target star-button starred upwards"
      title="Unstar this repository" data-remote="true" data-method="post" rel="nofollow">
      <span class="octicon octicon-star-delete"></span><span class="text">Unstar</span>
    </a>

    <a href="/klasbo/TTK4145/star"
      class="minibutton with-count js-toggler-target star-button unstarred upwards"
      title="Star this repository" data-remote="true" data-method="post" rel="nofollow">
      <span class="octicon octicon-star"></span><span class="text">Star</span>
    </a>

      <a class="social-count js-social-count" href="/klasbo/TTK4145/stargazers">
        3
      </a>
  </div>

  </li>


        <li>
          <a href="/klasbo/TTK4145/fork" class="minibutton with-count js-toggler-target fork-button lighter upwards" title="Fork this repo" rel="facebox nofollow">
            <span class="octicon octicon-git-branch-create"></span><span class="text">Fork</span>
          </a>
          <a href="/klasbo/TTK4145/network" class="social-count">5</a>
        </li>


</ul>

        <h1 itemscope itemtype="http://data-vocabulary.org/Breadcrumb" class="entry-title public">
          <span class="repo-label"><span>public</span></span>
          <span class="mega-octicon octicon-repo"></span>
          <span class="author">
            <a href="/klasbo" class="url fn" itemprop="url" rel="author"><span itemprop="title">klasbo</span></a>
          </span>
          <span class="repohead-name-divider">/</span>
          <strong><a href="/klasbo/TTK4145" class="js-current-repository js-repo-home-link">TTK4145</a></strong>

          <span class="page-context-loader">
            <img alt="Octocat-spinner-32" height="16" src="https://github.global.ssl.fastly.net/images/spinners/octocat-spinner-32.gif" width="16" />
          </span>

        </h1>
      </div><!-- /.container -->
    </div><!-- /.repohead -->

    <div class="container">
      

      <div class="repository-with-sidebar repo-container new-discussion-timeline js-new-discussion-timeline  ">
        <div class="repository-sidebar">
            

<div class="sunken-menu vertical-right repo-nav js-repo-nav js-repository-container-pjax js-octicon-loaders">
  <div class="sunken-menu-contents">
    <ul class="sunken-menu-group">
      <li class="tooltipped leftwards" title="Code">
        <a href="/klasbo/TTK4145" aria-label="Code" class="selected js-selected-navigation-item sunken-menu-item" data-gotokey="c" data-pjax="true" data-selected-links="repo_source repo_downloads repo_commits repo_tags repo_branches /klasbo/TTK4145">
          <span class="octicon octicon-code"></span> <span class="full-word">Code</span>
          <img alt="Octocat-spinner-32" class="mini-loader" height="16" src="https://github.global.ssl.fastly.net/images/spinners/octocat-spinner-32.gif" width="16" />
</a>      </li>

        <li class="tooltipped leftwards" title="Issues">
          <a href="/klasbo/TTK4145/issues" aria-label="Issues" class="js-selected-navigation-item sunken-menu-item js-disable-pjax" data-gotokey="i" data-selected-links="repo_issues /klasbo/TTK4145/issues">
            <span class="octicon octicon-issue-opened"></span> <span class="full-word">Issues</span>
            <span class='counter'>0</span>
            <img alt="Octocat-spinner-32" class="mini-loader" height="16" src="https://github.global.ssl.fastly.net/images/spinners/octocat-spinner-32.gif" width="16" />
</a>        </li>

      <li class="tooltipped leftwards" title="Pull Requests">
        <a href="/klasbo/TTK4145/pulls" aria-label="Pull Requests" class="js-selected-navigation-item sunken-menu-item js-disable-pjax" data-gotokey="p" data-selected-links="repo_pulls /klasbo/TTK4145/pulls">
            <span class="octicon octicon-git-pull-request"></span> <span class="full-word">Pull Requests</span>
            <span class='counter'>0</span>
            <img alt="Octocat-spinner-32" class="mini-loader" height="16" src="https://github.global.ssl.fastly.net/images/spinners/octocat-spinner-32.gif" width="16" />
</a>      </li>


        <li class="tooltipped leftwards" title="Wiki">
          <a href="/klasbo/TTK4145/wiki" aria-label="Wiki" class="js-selected-navigation-item sunken-menu-item" data-pjax="true" data-selected-links="repo_wiki /klasbo/TTK4145/wiki">
            <span class="octicon octicon-book"></span> <span class="full-word">Wiki</span>
            <img alt="Octocat-spinner-32" class="mini-loader" height="16" src="https://github.global.ssl.fastly.net/images/spinners/octocat-spinner-32.gif" width="16" />
</a>        </li>
    </ul>
    <div class="sunken-menu-separator"></div>
    <ul class="sunken-menu-group">

      <li class="tooltipped leftwards" title="Pulse">
        <a href="/klasbo/TTK4145/pulse" aria-label="Pulse" class="js-selected-navigation-item sunken-menu-item" data-pjax="true" data-selected-links="pulse /klasbo/TTK4145/pulse">
          <span class="octicon octicon-pulse"></span> <span class="full-word">Pulse</span>
          <img alt="Octocat-spinner-32" class="mini-loader" height="16" src="https://github.global.ssl.fastly.net/images/spinners/octocat-spinner-32.gif" width="16" />
</a>      </li>

      <li class="tooltipped leftwards" title="Graphs">
        <a href="/klasbo/TTK4145/graphs" aria-label="Graphs" class="js-selected-navigation-item sunken-menu-item" data-pjax="true" data-selected-links="repo_graphs repo_contributors /klasbo/TTK4145/graphs">
          <span class="octicon octicon-graph"></span> <span class="full-word">Graphs</span>
          <img alt="Octocat-spinner-32" class="mini-loader" height="16" src="https://github.global.ssl.fastly.net/images/spinners/octocat-spinner-32.gif" width="16" />
</a>      </li>

      <li class="tooltipped leftwards" title="Network">
        <a href="/klasbo/TTK4145/network" aria-label="Network" class="js-selected-navigation-item sunken-menu-item js-disable-pjax" data-selected-links="repo_network /klasbo/TTK4145/network">
          <span class="octicon octicon-git-branch"></span> <span class="full-word">Network</span>
          <img alt="Octocat-spinner-32" class="mini-loader" height="16" src="https://github.global.ssl.fastly.net/images/spinners/octocat-spinner-32.gif" width="16" />
</a>      </li>
    </ul>


  </div>
</div>

              <div class="only-with-full-nav">
                

  

<div class="clone-url open"
  data-protocol-type="http"
  data-url="/users/set_protocol?protocol_selector=http&amp;protocol_type=clone">
  <h3><strong>HTTPS</strong> clone URL</h3>
  <div class="clone-url-box">
    <input type="text" class="clone js-url-field"
           value="https://github.com/klasbo/TTK4145.git" readonly="readonly">

    <span class="js-zeroclipboard url-box-clippy minibutton zeroclipboard-button" data-clipboard-text="https://github.com/klasbo/TTK4145.git" data-copied-hint="copied!" title="copy to clipboard"><span class="octicon octicon-clippy"></span></span>
  </div>
</div>

  

<div class="clone-url "
  data-protocol-type="ssh"
  data-url="/users/set_protocol?protocol_selector=ssh&amp;protocol_type=clone">
  <h3><strong>SSH</strong> clone URL</h3>
  <div class="clone-url-box">
    <input type="text" class="clone js-url-field"
           value="git@github.com:klasbo/TTK4145.git" readonly="readonly">

    <span class="js-zeroclipboard url-box-clippy minibutton zeroclipboard-button" data-clipboard-text="git@github.com:klasbo/TTK4145.git" data-copied-hint="copied!" title="copy to clipboard"><span class="octicon octicon-clippy"></span></span>
  </div>
</div>

  

<div class="clone-url "
  data-protocol-type="subversion"
  data-url="/users/set_protocol?protocol_selector=subversion&amp;protocol_type=clone">
  <h3><strong>Subversion</strong> checkout URL</h3>
  <div class="clone-url-box">
    <input type="text" class="clone js-url-field"
           value="https://github.com/klasbo/TTK4145" readonly="readonly">

    <span class="js-zeroclipboard url-box-clippy minibutton zeroclipboard-button" data-clipboard-text="https://github.com/klasbo/TTK4145" data-copied-hint="copied!" title="copy to clipboard"><span class="octicon octicon-clippy"></span></span>
  </div>
</div>


<p class="clone-options">You can clone with
      <a href="#" class="js-clone-selector" data-protocol="http">HTTPS</a>,
      <a href="#" class="js-clone-selector" data-protocol="ssh">SSH</a>,
      or <a href="#" class="js-clone-selector" data-protocol="subversion">Subversion</a>.
  <span class="octicon help tooltipped upwards" title="Get help on which URL is right for you.">
    <a href="https://help.github.com/articles/which-remote-url-should-i-use">
    <span class="octicon octicon-question"></span>
    </a>
  </span>
</p>



                <a href="/klasbo/TTK4145/archive/master.zip"
                   class="minibutton sidebar-button"
                   title="Download this repository as a zip file"
                   rel="nofollow">
                  <span class="octicon octicon-cloud-download"></span>
                  Download ZIP
                </a>
              </div>
        </div><!-- /.repository-sidebar -->

        <div id="js-repo-pjax-container" class="repository-content context-loader-container" data-pjax-container>
          


<!-- blob contrib key: blob_contributors:v21:0f8aa2b14c3ef20a45a1540e7fd9e7c0 -->

<p title="This is a placeholder element" class="js-history-link-replace hidden"></p>

<a href="/klasbo/TTK4145/find/master" data-pjax data-hotkey="t" class="js-show-file-finder" style="display:none">Show File Finder</a>

<div class="file-navigation">
  

<div class="select-menu js-menu-container js-select-menu" >
  <span class="minibutton select-menu-button js-menu-target" data-hotkey="w"
    data-master-branch="master"
    data-ref="master"
    role="button" aria-label="Switch branches or tags" tabindex="0">
    <span class="octicon octicon-git-branch"></span>
    <i>branch:</i>
    <span class="js-select-button">master</span>
  </span>

  <div class="select-menu-modal-holder js-menu-content js-navigation-container" data-pjax>

    <div class="select-menu-modal">
      <div class="select-menu-header">
        <span class="select-menu-title">Switch branches/tags</span>
        <span class="octicon octicon-remove-close js-menu-close"></span>
      </div> <!-- /.select-menu-header -->

      <div class="select-menu-filters">
        <div class="select-menu-text-filter">
          <input type="text" aria-label="Filter branches/tags" id="context-commitish-filter-field" class="js-filterable-field js-navigation-enable" placeholder="Filter branches/tags">
        </div>
        <div class="select-menu-tabs">
          <ul>
            <li class="select-menu-tab">
              <a href="#" data-tab-filter="branches" class="js-select-menu-tab">Branches</a>
            </li>
            <li class="select-menu-tab">
              <a href="#" data-tab-filter="tags" class="js-select-menu-tab">Tags</a>
            </li>
          </ul>
        </div><!-- /.select-menu-tabs -->
      </div><!-- /.select-menu-filters -->

      <div class="select-menu-list select-menu-tab-bucket js-select-menu-tab-bucket" data-tab-filter="branches">

        <div data-filterable-for="context-commitish-filter-field" data-filterable-type="substring">


            <div class="select-menu-item js-navigation-item selected">
              <span class="select-menu-item-icon octicon octicon-check"></span>
              <a href="/klasbo/TTK4145/blob/master/Project/driver/elev.h"
                 data-name="master"
                 data-skip-pjax="true"
                 rel="nofollow"
                 class="js-navigation-open select-menu-item-text js-select-button-text css-truncate-target"
                 title="master">master</a>
            </div> <!-- /.select-menu-item -->
        </div>

          <div class="select-menu-no-results">Nothing to show</div>
      </div> <!-- /.select-menu-list -->

      <div class="select-menu-list select-menu-tab-bucket js-select-menu-tab-bucket" data-tab-filter="tags">
        <div data-filterable-for="context-commitish-filter-field" data-filterable-type="substring">


        </div>

        <div class="select-menu-no-results">Nothing to show</div>
      </div> <!-- /.select-menu-list -->

    </div> <!-- /.select-menu-modal -->
  </div> <!-- /.select-menu-modal-holder -->
</div> <!-- /.select-menu -->

  <div class="breadcrumb">
    <span class='repo-root js-repo-root'><span itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/klasbo/TTK4145" data-branch="master" data-direction="back" data-pjax="true" itemscope="url"><span itemprop="title">TTK4145</span></a></span></span><span class="separator"> / </span><span itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/klasbo/TTK4145/tree/master/Project" data-branch="master" data-direction="back" data-pjax="true" itemscope="url"><span itemprop="title">Project</span></a></span><span class="separator"> / </span><span itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/klasbo/TTK4145/tree/master/Project/driver" data-branch="master" data-direction="back" data-pjax="true" itemscope="url"><span itemprop="title">driver</span></a></span><span class="separator"> / </span><strong class="final-path">elev.h</strong> <span class="js-zeroclipboard minibutton zeroclipboard-button" data-clipboard-text="Project/driver/elev.h" data-copied-hint="copied!" title="copy to clipboard"><span class="octicon octicon-clippy"></span></span>
  </div>
</div>


  <div class="commit file-history-tease">
    <img alt="FWorren" class="main-avatar" height="24" src="https://1.gravatar.com/avatar/2c6f220a5887389e312a48085334a468?d=https%3A%2F%2Fidenticons.github.com%2Fe1a25ff1ee8d1df75858b4a83bfee064.png&amp;r=x&amp;s=140" width="24" />
    <span class="author"><a href="/FWorren" rel="author">FWorren</a></span>
    <time class="js-relative-date" data-title-format="YYYY-MM-DD HH:mm:ss" datetime="2014-02-06T09:18:31-08:00" title="2014-02-06 09:18:31">February 06, 2014</time>
    <div class="commit-title">
        <a href="/klasbo/TTK4145/commit/64551ef1eca244838f734d143e8e8b40d256e88d" class="message" data-pjax="true" title="Added cgo stuff">Added cgo stuff</a>
    </div>

    <div class="participation">
      <p class="quickstat"><a href="#blob_contributors_box" rel="facebox"><strong>2</strong> contributors</a></p>
          <a class="avatar tooltipped downwards" title="klasbo" href="/klasbo/TTK4145/commits/master/Project/driver/elev.h?author=klasbo"><img alt="klasbo" height="20" src="https://0.gravatar.com/avatar/15cebd3d1dfcbf51093ec71bb5649e2a?d=https%3A%2F%2Fidenticons.github.com%2Ff4f9cec9fec60d40d51bfdbb5894cc36.png&amp;r=x&amp;s=140" width="20" /></a>
    <a class="avatar tooltipped downwards" title="FWorren" href="/klasbo/TTK4145/commits/master/Project/driver/elev.h?author=FWorren"><img alt="FWorren" height="20" src="https://1.gravatar.com/avatar/2c6f220a5887389e312a48085334a468?d=https%3A%2F%2Fidenticons.github.com%2Fe1a25ff1ee8d1df75858b4a83bfee064.png&amp;r=x&amp;s=140" width="20" /></a>


    </div>
    <div id="blob_contributors_box" style="display:none">
      <h2 class="facebox-header">Users who have contributed to this file</h2>
      <ul class="facebox-user-list">
          <li class="facebox-user-list-item">
            <img alt="klasbo" height="24" src="https://0.gravatar.com/avatar/15cebd3d1dfcbf51093ec71bb5649e2a?d=https%3A%2F%2Fidenticons.github.com%2Ff4f9cec9fec60d40d51bfdbb5894cc36.png&amp;r=x&amp;s=140" width="24" />
            <a href="/klasbo">klasbo</a>
          </li>
          <li class="facebox-user-list-item">
            <img alt="FWorren" height="24" src="https://1.gravatar.com/avatar/2c6f220a5887389e312a48085334a468?d=https%3A%2F%2Fidenticons.github.com%2Fe1a25ff1ee8d1df75858b4a83bfee064.png&amp;r=x&amp;s=140" width="24" />
            <a href="/FWorren">FWorren</a>
          </li>
      </ul>
    </div>
  </div>

<div id="files" class="bubble">
  <div class="file">
    <div class="meta">
      <div class="info">
        <span class="icon"><b class="octicon octicon-file-text"></b></span>
        <span class="mode" title="File Mode">file</span>
          <span>114 lines (76 sloc)</span>
        <span>2.653 kb</span>
      </div>
      <div class="actions">
        <div class="button-group">
                <a class="minibutton tooltipped upwards js-update-url-with-hash"
                   title="Clicking this button will automatically fork this project so you can edit the file"
                   href="/klasbo/TTK4145/edit/master/Project/driver/elev.h"
                   data-method="post" rel="nofollow">Edit</a>
          <a href="/klasbo/TTK4145/raw/master/Project/driver/elev.h" class="button minibutton " id="raw-url">Raw</a>
            <a href="/klasbo/TTK4145/blame/master/Project/driver/elev.h" class="button minibutton js-update-url-with-hash">Blame</a>
          <a href="/klasbo/TTK4145/commits/master/Project/driver/elev.h" class="button minibutton " rel="nofollow">History</a>
        </div><!-- /.button-group -->
          <a class="minibutton danger empty-icon tooltipped downwards"
             href="/klasbo/TTK4145/delete/master/Project/driver/elev.h"
             title="Fork this project and delete file"
             data-method="post" data-test-id="delete-blob-file" rel="nofollow">
          Delete
        </a>
      </div><!-- /.actions -->
    </div>
        <div class="blob-wrapper data type-c js-blob-data">
        <table class="file-code file-diff tab-size-8">
          <tr class="file-code-line">
            <td class="blob-line-nums">
              <span id="L1" rel="#L1">1</span>
<span id="L2" rel="#L2">2</span>
<span id="L3" rel="#L3">3</span>
<span id="L4" rel="#L4">4</span>
<span id="L5" rel="#L5">5</span>
<span id="L6" rel="#L6">6</span>
<span id="L7" rel="#L7">7</span>
<span id="L8" rel="#L8">8</span>
<span id="L9" rel="#L9">9</span>
<span id="L10" rel="#L10">10</span>
<span id="L11" rel="#L11">11</span>
<span id="L12" rel="#L12">12</span>
<span id="L13" rel="#L13">13</span>
<span id="L14" rel="#L14">14</span>
<span id="L15" rel="#L15">15</span>
<span id="L16" rel="#L16">16</span>
<span id="L17" rel="#L17">17</span>
<span id="L18" rel="#L18">18</span>
<span id="L19" rel="#L19">19</span>
<span id="L20" rel="#L20">20</span>
<span id="L21" rel="#L21">21</span>
<span id="L22" rel="#L22">22</span>
<span id="L23" rel="#L23">23</span>
<span id="L24" rel="#L24">24</span>
<span id="L25" rel="#L25">25</span>
<span id="L26" rel="#L26">26</span>
<span id="L27" rel="#L27">27</span>
<span id="L28" rel="#L28">28</span>
<span id="L29" rel="#L29">29</span>
<span id="L30" rel="#L30">30</span>
<span id="L31" rel="#L31">31</span>
<span id="L32" rel="#L32">32</span>
<span id="L33" rel="#L33">33</span>
<span id="L34" rel="#L34">34</span>
<span id="L35" rel="#L35">35</span>
<span id="L36" rel="#L36">36</span>
<span id="L37" rel="#L37">37</span>
<span id="L38" rel="#L38">38</span>
<span id="L39" rel="#L39">39</span>
<span id="L40" rel="#L40">40</span>
<span id="L41" rel="#L41">41</span>
<span id="L42" rel="#L42">42</span>
<span id="L43" rel="#L43">43</span>
<span id="L44" rel="#L44">44</span>
<span id="L45" rel="#L45">45</span>
<span id="L46" rel="#L46">46</span>
<span id="L47" rel="#L47">47</span>
<span id="L48" rel="#L48">48</span>
<span id="L49" rel="#L49">49</span>
<span id="L50" rel="#L50">50</span>
<span id="L51" rel="#L51">51</span>
<span id="L52" rel="#L52">52</span>
<span id="L53" rel="#L53">53</span>
<span id="L54" rel="#L54">54</span>
<span id="L55" rel="#L55">55</span>
<span id="L56" rel="#L56">56</span>
<span id="L57" rel="#L57">57</span>
<span id="L58" rel="#L58">58</span>
<span id="L59" rel="#L59">59</span>
<span id="L60" rel="#L60">60</span>
<span id="L61" rel="#L61">61</span>
<span id="L62" rel="#L62">62</span>
<span id="L63" rel="#L63">63</span>
<span id="L64" rel="#L64">64</span>
<span id="L65" rel="#L65">65</span>
<span id="L66" rel="#L66">66</span>
<span id="L67" rel="#L67">67</span>
<span id="L68" rel="#L68">68</span>
<span id="L69" rel="#L69">69</span>
<span id="L70" rel="#L70">70</span>
<span id="L71" rel="#L71">71</span>
<span id="L72" rel="#L72">72</span>
<span id="L73" rel="#L73">73</span>
<span id="L74" rel="#L74">74</span>
<span id="L75" rel="#L75">75</span>
<span id="L76" rel="#L76">76</span>
<span id="L77" rel="#L77">77</span>
<span id="L78" rel="#L78">78</span>
<span id="L79" rel="#L79">79</span>
<span id="L80" rel="#L80">80</span>
<span id="L81" rel="#L81">81</span>
<span id="L82" rel="#L82">82</span>
<span id="L83" rel="#L83">83</span>
<span id="L84" rel="#L84">84</span>
<span id="L85" rel="#L85">85</span>
<span id="L86" rel="#L86">86</span>
<span id="L87" rel="#L87">87</span>
<span id="L88" rel="#L88">88</span>
<span id="L89" rel="#L89">89</span>
<span id="L90" rel="#L90">90</span>
<span id="L91" rel="#L91">91</span>
<span id="L92" rel="#L92">92</span>
<span id="L93" rel="#L93">93</span>
<span id="L94" rel="#L94">94</span>
<span id="L95" rel="#L95">95</span>
<span id="L96" rel="#L96">96</span>
<span id="L97" rel="#L97">97</span>
<span id="L98" rel="#L98">98</span>
<span id="L99" rel="#L99">99</span>
<span id="L100" rel="#L100">100</span>
<span id="L101" rel="#L101">101</span>
<span id="L102" rel="#L102">102</span>
<span id="L103" rel="#L103">103</span>
<span id="L104" rel="#L104">104</span>
<span id="L105" rel="#L105">105</span>
<span id="L106" rel="#L106">106</span>
<span id="L107" rel="#L107">107</span>
<span id="L108" rel="#L108">108</span>
<span id="L109" rel="#L109">109</span>
<span id="L110" rel="#L110">110</span>
<span id="L111" rel="#L111">111</span>
<span id="L112" rel="#L112">112</span>
<span id="L113" rel="#L113">113</span>

            </td>
            <td class="blob-line-code"><div class="code-body highlight"><pre><div class='line' id='LC1'><span class="c1">// Wrapper for libComedi Elevator control.</span></div><div class='line' id='LC2'><span class="c1">// These functions provides an interface to the elevators in the real time lab</span></div><div class='line' id='LC3'><span class="c1">//</span></div><div class='line' id='LC4'><span class="c1">// 2007, Martin Korsgaard</span></div><div class='line' id='LC5'><span class="cp">#ifndef __INCLUDE_DRIVER_H__</span></div><div class='line' id='LC6'><span class="cp">#define __INCLUDE_DRIVER_H__</span></div><div class='line' id='LC7'><br/></div><div class='line' id='LC8'><br/></div><div class='line' id='LC9'><span class="c1">// Number of floors</span></div><div class='line' id='LC10'><span class="cp">#define N_FLOORS 4</span></div><div class='line' id='LC11'><br/></div><div class='line' id='LC12'><span class="cm">/**</span></div><div class='line' id='LC13'><span class="cm">  Button types for function elev_set_button_lamp() and elev_get_button().</span></div><div class='line' id='LC14'><span class="cm">*/</span></div><div class='line' id='LC15'><span class="k">typedef</span> <span class="k">enum</span> <span class="n">tag_elev_lamp_type</span> <span class="p">{</span> </div><div class='line' id='LC16'>&nbsp;&nbsp;&nbsp;&nbsp;<span class="n">BUTTON_CALL_UP</span> <span class="o">=</span> <span class="mi">0</span><span class="p">,</span> </div><div class='line' id='LC17'>&nbsp;&nbsp;&nbsp;&nbsp;<span class="n">BUTTON_CALL_DOWN</span> <span class="o">=</span> <span class="mi">1</span><span class="p">,</span> </div><div class='line' id='LC18'>&nbsp;&nbsp;&nbsp;&nbsp;<span class="n">BUTTON_COMMAND</span> <span class="o">=</span> <span class="mi">2</span> </div><div class='line' id='LC19'><span class="p">}</span> <span class="kt">elev_button_type_t</span><span class="p">;</span></div><div class='line' id='LC20'><br/></div><div class='line' id='LC21'><br/></div><div class='line' id='LC22'><br/></div><div class='line' id='LC23'><span class="cm">/**</span></div><div class='line' id='LC24'><span class="cm">  Initialize elevator.</span></div><div class='line' id='LC25'><span class="cm">  @return Non-zero on success, 0 on failure.</span></div><div class='line' id='LC26'><span class="cm">*/</span></div><div class='line' id='LC27'><span class="kt">int</span> <span class="nf">elev_init</span><span class="p">(</span><span class="kt">void</span><span class="p">);</span></div><div class='line' id='LC28'><br/></div><div class='line' id='LC29'><br/></div><div class='line' id='LC30'><br/></div><div class='line' id='LC31'><span class="cm">/**</span></div><div class='line' id='LC32'><span class="cm">  Sets the speed of the elevator. </span></div><div class='line' id='LC33'><span class="cm">  @param speed New speed of elevator. Positive values denote upward movement</span></div><div class='line' id='LC34'><span class="cm">    and vice versa. Set speed to 0 to stop the elevator. From -300 to 300 gives</span></div><div class='line' id='LC35'><span class="cm">    sensible speeds. The hardware wil emit a constant tone if the speed is too high.</span></div><div class='line' id='LC36'><span class="cm">*/</span></div><div class='line' id='LC37'><span class="kt">void</span> <span class="nf">elev_set_speed</span><span class="p">(</span><span class="kt">int</span> <span class="n">speed</span><span class="p">);</span></div><div class='line' id='LC38'><br/></div><div class='line' id='LC39'><br/></div><div class='line' id='LC40'><br/></div><div class='line' id='LC41'><span class="cm">/**</span></div><div class='line' id='LC42'><span class="cm">  Get floor sensor signal.</span></div><div class='line' id='LC43'><span class="cm">  @return -1 if elevator is not on a floor. 0-3 if elevator is on floor. 0 is</span></div><div class='line' id='LC44'><span class="cm">    ground floor, 3 is top floor.</span></div><div class='line' id='LC45'><span class="cm">*/</span></div><div class='line' id='LC46'><span class="kt">int</span> <span class="nf">elev_get_floor_sensor_signal</span><span class="p">(</span><span class="kt">void</span><span class="p">);</span></div><div class='line' id='LC47'><br/></div><div class='line' id='LC48'><br/></div><div class='line' id='LC49'><br/></div><div class='line' id='LC50'><span class="cm">/**</span></div><div class='line' id='LC51'><span class="cm">  Gets a button signal.</span></div><div class='line' id='LC52'><span class="cm">  @param button Which button type to check. Can be BUTTON_CALL_UP,</span></div><div class='line' id='LC53'><span class="cm">    BUTTON_CALL_DOWN or BUTTON_COMMAND (button &quot;inside the elevator).</span></div><div class='line' id='LC54'><span class="cm">  @param floor Which floor to check button. Must be 0-3.</span></div><div class='line' id='LC55'><span class="cm">  @return 0 if button is not pushed. 1 if button is pushed.</span></div><div class='line' id='LC56'><span class="cm">*/</span></div><div class='line' id='LC57'><span class="kt">int</span> <span class="nf">elev_get_button_signal</span><span class="p">(</span><span class="kt">elev_button_type_t</span> <span class="n">button</span><span class="p">,</span> <span class="kt">int</span> <span class="n">floor</span><span class="p">);</span></div><div class='line' id='LC58'><br/></div><div class='line' id='LC59'><br/></div><div class='line' id='LC60'><br/></div><div class='line' id='LC61'><span class="cm">/**</span></div><div class='line' id='LC62'><span class="cm">  Get signal from stop button.</span></div><div class='line' id='LC63'><span class="cm">  @return 1 if stop button is pushed, 0 if not.</span></div><div class='line' id='LC64'><span class="cm">*/</span></div><div class='line' id='LC65'><span class="kt">int</span> <span class="nf">elev_get_stop_signal</span><span class="p">(</span><span class="kt">void</span><span class="p">);</span></div><div class='line' id='LC66'><br/></div><div class='line' id='LC67'><br/></div><div class='line' id='LC68'><br/></div><div class='line' id='LC69'><span class="cm">/**</span></div><div class='line' id='LC70'><span class="cm">  Get signal from obstruction switch.</span></div><div class='line' id='LC71'><span class="cm">  @return 1 if obstruction is enabled. 0 if not.</span></div><div class='line' id='LC72'><span class="cm">*/</span></div><div class='line' id='LC73'><span class="kt">int</span> <span class="nf">elev_get_obstruction_signal</span><span class="p">(</span><span class="kt">void</span><span class="p">);</span></div><div class='line' id='LC74'><br/></div><div class='line' id='LC75'><br/></div><div class='line' id='LC76'><br/></div><div class='line' id='LC77'><span class="cm">/**</span></div><div class='line' id='LC78'><span class="cm">  Set floor indicator lamp for a given floor.</span></div><div class='line' id='LC79'><span class="cm">  @param floor Which floor lamp to turn on. Other floor lamps are turned off.</span></div><div class='line' id='LC80'><span class="cm">*/</span></div><div class='line' id='LC81'><span class="kt">void</span> <span class="nf">elev_set_floor_indicator</span><span class="p">(</span><span class="kt">int</span> <span class="n">floor</span><span class="p">);</span></div><div class='line' id='LC82'><br/></div><div class='line' id='LC83'><br/></div><div class='line' id='LC84'><br/></div><div class='line' id='LC85'><span class="cm">/**</span></div><div class='line' id='LC86'><span class="cm">  Set a button lamp.</span></div><div class='line' id='LC87'><span class="cm">  @param lamp Which type of lamp to set. Can be BUTTON_CALL_UP,</span></div><div class='line' id='LC88'><span class="cm">    BUTTON_CALL_DOWN or BUTTON_COMMAND (button &quot;inside&quot; the elevator).</span></div><div class='line' id='LC89'><span class="cm">  @param floor Floor of lamp to set. Must be 0-3</span></div><div class='line' id='LC90'><span class="cm">  @param value Non-zero value turns lamp on, 0 turns lamp off.</span></div><div class='line' id='LC91'><span class="cm">*/</span></div><div class='line' id='LC92'><span class="kt">void</span> <span class="nf">elev_set_button_lamp</span><span class="p">(</span><span class="kt">elev_button_type_t</span> <span class="n">button</span><span class="p">,</span> <span class="kt">int</span> <span class="n">floor</span><span class="p">,</span> <span class="kt">int</span> <span class="n">value</span><span class="p">);</span></div><div class='line' id='LC93'><br/></div><div class='line' id='LC94'><br/></div><div class='line' id='LC95'><br/></div><div class='line' id='LC96'><span class="cm">/**</span></div><div class='line' id='LC97'><span class="cm">  Turn stop lamp on or off.</span></div><div class='line' id='LC98'><span class="cm">  @param value Non-zero value turns lamp on, 0 turns lamp off.</span></div><div class='line' id='LC99'><span class="cm">*/</span></div><div class='line' id='LC100'><span class="kt">void</span> <span class="nf">elev_set_stop_lamp</span><span class="p">(</span><span class="kt">int</span> <span class="n">value</span><span class="p">);</span></div><div class='line' id='LC101'><br/></div><div class='line' id='LC102'><br/></div><div class='line' id='LC103'><br/></div><div class='line' id='LC104'><span class="cm">/**</span></div><div class='line' id='LC105'><span class="cm">  Turn door-open lamp on or off.</span></div><div class='line' id='LC106'><span class="cm">  @param value Non-zero value turns lamp on, 0 turns lamp off.</span></div><div class='line' id='LC107'><span class="cm">*/</span></div><div class='line' id='LC108'><span class="kt">void</span> <span class="nf">elev_set_door_open_lamp</span><span class="p">(</span><span class="kt">int</span> <span class="n">value</span><span class="p">);</span></div><div class='line' id='LC109'><br/></div><div class='line' id='LC110'><br/></div><div class='line' id='LC111'><br/></div><div class='line' id='LC112'><span class="cp">#endif </span><span class="c1">// #ifndef __INCLUDE_DRIVER_H__</span></div><div class='line' id='LC113'><br/></div></pre></div></td>
          </tr>
        </table>
  </div>

  </div>
</div>

<a href="#jump-to-line" rel="facebox[.linejump]" data-hotkey="l" class="js-jump-to-line" style="display:none">Jump to Line</a>
<div id="jump-to-line" style="display:none">
  <form accept-charset="UTF-8" class="js-jump-to-line-form">
    <input class="linejump-input js-jump-to-line-field" type="text" placeholder="Jump to line&hellip;" autofocus>
    <button type="submit" class="button">Go</button>
  </form>
</div>

        </div>

      </div><!-- /.repo-container -->
      <div class="modal-backdrop"></div>
    </div><!-- /.container -->
  </div><!-- /.site -->


    </div><!-- /.wrapper -->

      <div class="container">
  <div class="site-footer">
    <ul class="site-footer-links right">
      <li><a href="https://status.github.com/">Status</a></li>
      <li><a href="http://developer.github.com">API</a></li>
      <li><a href="http://training.github.com">Training</a></li>
      <li><a href="http://shop.github.com">Shop</a></li>
      <li><a href="/blog">Blog</a></li>
      <li><a href="/about">About</a></li>

    </ul>

    <a href="/">
      <span class="mega-octicon octicon-mark-github" title="GitHub"></span>
    </a>

    <ul class="site-footer-links">
      <li>&copy; 2014 <span title="0.04085s from github-fe132-cp1-prd.iad.github.net">GitHub</span>, Inc.</li>
        <li><a href="/site/terms">Terms</a></li>
        <li><a href="/site/privacy">Privacy</a></li>
        <li><a href="/security">Security</a></li>
        <li><a href="/contact">Contact</a></li>
    </ul>
  </div><!-- /.site-footer -->
</div><!-- /.container -->


    <div class="fullscreen-overlay js-fullscreen-overlay" id="fullscreen_overlay">
  <div class="fullscreen-container js-fullscreen-container">
    <div class="textarea-wrap">
      <textarea name="fullscreen-contents" id="fullscreen-contents" class="js-fullscreen-contents" placeholder="" data-suggester="fullscreen_suggester"></textarea>
          <div class="suggester-container">
              <div class="suggester fullscreen-suggester js-navigation-container" id="fullscreen_suggester"
                 data-url="/klasbo/TTK4145/suggestions/commit">
              </div>
          </div>
    </div>
  </div>
  <div class="fullscreen-sidebar">
    <a href="#" class="exit-fullscreen js-exit-fullscreen tooltipped leftwards" title="Exit Zen Mode">
      <span class="mega-octicon octicon-screen-normal"></span>
    </a>
    <a href="#" class="theme-switcher js-theme-switcher tooltipped leftwards"
      title="Switch themes">
      <span class="octicon octicon-color-mode"></span>
    </a>
  </div>
</div>



    <div id="ajax-error-message" class="flash flash-error">
      <span class="octicon octicon-alert"></span>
      <a href="#" class="octicon octicon-remove-close close js-ajax-error-dismiss"></a>
      Something went wrong with that request. Please try again.
    </div>

  </body>
</html>

