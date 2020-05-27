import csv
import textwrap
from datetime import datetime, timedelta
import pytz
from tzlocal import get_localzone


class Airport:
    """ Airport as represented in airports.dat """
    # CSV columns will become Airport property names in __init__
    prop_names = ('id', 'name', 'city', 'country', 'iata', 'icao', 'lat', 'long', 'alt', 'utc_offset', 'dst_rule',
                  'tz', 'type', 'source')

    def __init__(self, csv_entry):
        assert len(csv_entry) == len(Airport.prop_names)
        self.__dict__.update(dict(zip(Airport.prop_names, csv_entry)))

    def __str__(self):
        return "{0.iata} ({0.name})".format(self)


def load_airports(csv_file_name):
    """ Load airports into a dictionary where keys are IATA codes"""
    airports = {}
    with open(csv_file_name, newline='') as data_file:
        for entry in csv.reader(data_file):
            a = Airport(csv_entry=entry)
            airports[a.iata] = a
    return airports


class Flight:
    def __init__(self, flight_id, origin, destination, departure, arrival):
        self.id = flight_id
        self.origin = origin
        self.destination = destination

        self.departure = self.localize_flight_datetime(departure, origin.tz)
        self.arrival = self.localize_flight_datetime(arrival, destination.tz)

    @staticmethod
    def localize_flight_datetime(date_time, tz_name):
        tz = pytz.timezone(tz_name)
        try:
            return tz.localize(date_time, is_dst=None)
        except pytz.exceptions.AmbiguousTimeError:
            return tz.localize(date_time, is_dst=True)

    @property
    def check_in(self):
        return self.departure.tzinfo.normalize(self.departure - timedelta(hours=3))

    @property
    def duration(self):
        return self.arrival - self.departure

    def time_to_departure(self):
        return self.departure - get_localzone().localize(datetime.now())

    def __str__(self):
        return textwrap.dedent('''\
        Flight {0.id}:
            from        : {0.origin}
            to          : {0.destination}
            departure   : {0.departure} {0.departure.tzinfo}
            arrival     : {0.arrival} {0.arrival.tzinfo}
            duration    : {0.duration}

            time to departure       : {ttd}
            check-in                : {0.check_in}
        '''.format(self, ttd = self.time_to_departure()))


airports = load_airports('airports.dat')

flights = [
    Flight(flight_id='AA123',
           origin=airports['ATL'],
           destination=airports['SVO'],
           departure=datetime(2018, 1, 1, 10, 10, 0),
           arrival=datetime(2018, 1, 2, 7, 12, 0)),
    Flight(flight_id='DST01',
           origin=airports['BRU'],
           destination=airports['ATL'],
           departure=datetime(2018, 3, 25, 1, 10, 0),
           arrival=datetime(2018, 3, 25, 7, 0, 0)),
    Flight(flight_id='DST02',
           origin=airports['BRU'],
           destination=airports['ATL'],
           departure=datetime(2018, 10, 28, 2, 10, 0),
           arrival=datetime(2018, 3, 25, 7, 0, 0))
]

#for f in flights:
#    print(f)
