# SunGrid Engine Serialization

This library will allow you to take the SGE XML output from queue status and serialize it programmatically into Go objects. As a note, natively unserializing will provide a ResourceList object containing all of the dynamic resources serialized into strings.

That's not exactly sane for programming, so several helper / receiver methods have been added to the ReceiverList allowing you to get structured values back for the important ones. 