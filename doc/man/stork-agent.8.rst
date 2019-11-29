..
   Copyright (C) 2019 Internet Systems Consortium, Inc. ("ISC")

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.

   See the COPYRIGHT file distributed with this work for additional
   information regarding copyright ownership.


stork-agent - Stork agent that monitors BIND and Kea services
-------------------------------------------------------------

Synopsis
~~~~~~~~

:program:`stork-agent` 

Description
~~~~~~~~~~~

The ``stork-agent`` is a small tool that is being run on the systems
that are running BIND and Kea services. Stork server connects to
the stork agent and uses it to monitor services remotely.


Configuration
~~~~~~~~~~~~~

Stork agent uses two environment variables to control its behavior:

- STORK_AGENT_ADDRESS - if defined, governs which IP address to listen on

- STORK_AGENT_PORT - if defined, it controls which port to listen on. The
  default is 8080.


Mailing List and Support
~~~~~~~~~~~~~~~~~~~~~~~~~

There is a public mailing list available for the Stork project. **stork-dev**
(stork-dev at lists.isc.org) is intended for Kea developers, prospective
contributors, and other advanced users. The list is available at
https://lists.isc.org. The community provides best-effort support
on both of those lists.

Once stork will become more mature, ISC will be providing professional support
for Stork services.

History
~~~~~~~

The ``stork-agent`` was first coded in November 2019 by Michal Nowikowski.

See Also
~~~~~~~~

:manpage:`stork-server(8)`