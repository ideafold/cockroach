module Utils {
    export module Format {
        // This function was adapted from
        // https://stackoverflow.com/questions/10420352/converting-file-size-in-bytes-to-human-readable
        var kibi = 1024
        var units = ['KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB'];
        export function Bytes(bytes: number): string {
            if (Math.abs(bytes) < kibi) {
                return bytes + ' B';
            }
            var u = -1;
            do {
                bytes /= kibi;
                ++u;
            } while (Math.abs(bytes) >= kibi && u < units.length - 1);
            return bytes.toFixed(1) + ' ' + units[u];
        }
    }
}
