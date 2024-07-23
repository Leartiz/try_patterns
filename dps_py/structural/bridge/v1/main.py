import abc

# ------------------------------------------------------------------------

class Device(abc.ABC):
    @abc.abstractmethod
    def is_enabled() -> bool: pass

    @abc.abstractmethod
    def enable(): pass

    @abc.abstractmethod
    def disable(): pass

    @abc.abstractmethod
    def get_volume() -> float: pass

    @abc.abstractmethod
    def set_volume(volume: float): pass

    @abc.abstractmethod
    def get_channel() -> int: pass

    @abc.abstractmethod
    def set_channel(channel: int): pass


class Tv(Device):
    def __init__(self) -> None:
        super().__init__()
        self.enabled = False
        self.channel = 0
        self.volume = 0

    def is_enabled(self) -> bool:
        print("Tv.is_enabled")
        return self.enabled

    def enable(self):
        print(f"Tv.enable")
        self.enabled = True

    def disable(self):
        print(f"Tv.disable")
        self.enabled = False

    def get_volume(self) -> float:
        print(f"Tv.get_volume")
        return self.volume

    def set_volume(self, volume: float):
        print(f"Tv.set_volume({volume})")
        if volume < 0:
            self.volume = 0
        elif volume > 100:
            self.volume = 100
        else:
            self.volume = volume

    def get_channel(self) -> int:
        print(f"Tv.get_channel")
        return self.channel

    def set_channel(self, channel: int):
        print(f"Tv.set_channel({channel})")
        if self.channel < 0:
            self.channel = 0
        else:
            self.channel = channel        

# ------------------------------------------------------------------------

# Пульт управления устройством.
class Remote():
    device: Device
    
    def __init__(self, device: Device) -> None:
        self.device = device

    def toggle_power(self):
        if self.device.is_enabled():
            self.device.disable()
        else:
            self.device.enable()

    # ***
    
    def volume_down(self):
        self.device.set_volume(
            self.device.get_volume() - 10)

    def volume_up(self):
        self.device.set_volume(
            self.device.get_volume() + 10)

    def channel_down(self):
        self.device.set_channel(
            self.device.get_channel() - 1)

    def channel_up(self):
        self.device.set_channel(
            self.device.get_channel() + 1)


class AdvancedRemote(Remote):
    def __init__(self, device: Device) -> None:
        super().__init__(device)

    def mute(self):
        self.device.set_volume(0)

# ------------------------------------------------------------------------

tv     = Tv()
remote = Remote(tv)

remote.volume_up()
remote.volume_up()
remote.volume_up()
remote.volume_up()
remote.volume_up()
remote.volume_down()

print(f"tv.volume: {tv.volume}")
print(f"tv.channel: {tv.channel}")
print(f"tv.enabled: {tv.enabled}")

# ------------------------------------------------------------------------