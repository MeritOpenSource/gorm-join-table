Code to model checking in to events at kiosks.

Context:
A kiosk enables users to check in to one of several events. At the same time, multiple kiosks can reference the same
event. This many-to-many relationship is represented in the database as a join table.  
The join table links a kiosk id and event id, thereby adding the event to the kiosk.

A checkin for an event happens at a kiosk configured with that event. Therefore, the checkins live at the intersection
of an event and a kiosk.

This repo contains code to model this relationship using:

* golang
* gorm
* liquibase
