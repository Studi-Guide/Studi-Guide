using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using StudiGuideAppService;

namespace StudiGuideAppService.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class RoomCoordinateController : ControllerBase
    {
        private static readonly string[] Summaries = new[]
        {
            "N413", "N231", "N321", "N5309", "K21037", "K213", "N102", "K01", "A01", "M01"
        };

        private readonly ILogger<RoomCoordinateController> _logger;

        public RoomCoordinateController(ILogger<RoomCoordinateController> logger)
        {
            _logger = logger;
        }

        [HttpGet]
        public IEnumerable<OhmRoom> Get()
        {
            var rng = new Random();
            return Enumerable.Range(1, 5).Select(index => new OhmRoom
            {
                YCoordinate = rng.Next(-20, 55),
                XCoordinate = rng.Next(-20, 55),
                ZCoordinate = rng.Next(-20, 55),
                RoomName = Summaries[rng.Next(Summaries.Length)]
            })
            .ToArray();
        }
    }
}
